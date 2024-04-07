package transform

import (
	"fmt"

	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	"github.com/streamingfast/bstream/transform"
	"github.com/streamingfast/dstore"
	pbtransform "github.com/streamingfast/firehose-acme/pb/zklend/starknet/transform/v1"
	pbacme "github.com/streamingfast/firehose-acme/pb/zklend/starknet/type/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Address string
type Topic string
type TopicWithRanges map[Topic][]BlockRange

const FELT_BYTE_LENGTH = 32

var TransactionEventFilterMessageName = proto.MessageName(&pbtransform.TransactionEventFilter{})

func NewTransactionEventFilterTransformFactory(_ dstore.Store, _ []uint64) (*transform.Factory, error) {
	return TransactionEventFilterTransformFactory, nil
}

var TransactionEventFilterTransformFactory = &transform.Factory{
	Obj: &pbtransform.TransactionEventFilter{},
	NewFunc: func(message *anypb.Any) (transform.Transform, error) {
		mname := message.MessageName()
		if mname != TransactionEventFilterMessageName {
			return nil, fmt.Errorf("expected type url %q, recevied %q ", TransactionEventFilterMessageName, message.TypeUrl)
		}

		filter := &pbtransform.TransactionEventFilter{}
		err := proto.Unmarshal(message.Value, filter)
		if err != nil {
			return nil, fmt.Errorf("unexpected unmarshall error: %w", err)
		}

		if len(filter.EventFilters) == 0 {
			return nil, fmt.Errorf("event filter must not be empty")
		}

		eventFilters := map[Address]TopicWithRanges{}

		for _, addressAndTopics := range filter.EventFilters {
			if len(addressAndTopics.ContractAddress) != FELT_BYTE_LENGTH {
				return nil, fmt.Errorf("address must be %d bytes, received %d bytes", FELT_BYTE_LENGTH, len(addressAndTopics.ContractAddress))
			}

			topics := TopicWithRanges{}

			for _, topicAndRanges := range addressAndTopics.Topics {
				if len(topicAndRanges.Topic) != FELT_BYTE_LENGTH {
					return nil, fmt.Errorf("topic must be %d bytes, received %d bytes", FELT_BYTE_LENGTH, len(topicAndRanges.Topic))
				}

				ranges := []BlockRange{}

				for _, blockRange := range topicAndRanges.BlockRanges {
					ranges = append(ranges, BlockRange{
						startBlock: blockRange.StartBlock,
						endBlock:   blockRange.EndBlock,
					})
				}

				topics[Topic(topicAndRanges.Topic)] = ranges
			}

			eventFilters[Address(addressAndTopics.ContractAddress)] = topics
		}

		return &TransactionEventFilter{
			eventFiltersByAddressAndTopic: eventFilters,
		}, nil
	},
}

type TransactionEventFilter struct {
	eventFiltersByAddressAndTopic map[Address]TopicWithRanges
}
type BlockRange struct {
	startBlock uint64
	endBlock   uint64
}

func (p *TransactionEventFilter) String() string {
	return "event-based block transaction filter"
}

func (p *TransactionEventFilter) Transform(readOnlyBlk *pbbstream.Block, in transform.Input) (transform.Output, error) {
	fullBlock := &pbacme.Block{}
	err := readOnlyBlk.Payload.UnmarshalTo(fullBlock)
	if err != nil {
		return nil, fmt.Errorf("mashalling block: %w", err)
	}

	filteredBlock := pbacme.Block{
		Height:       fullBlock.Height,
		Hash:         fullBlock.Hash,
		PrevHash:     fullBlock.PrevHash,
		Timestamp:    fullBlock.Timestamp,
		Transactions: []*pbacme.Transaction{},
	}

	for _, transaction := range fullBlock.Transactions {
		if p.isTransactionMatched(fullBlock, transaction) {
			filteredBlock.Transactions = append(filteredBlock.Transactions, transaction)
		}
	}

	return &filteredBlock, nil
}

func (p *TransactionEventFilter) isTransactionMatched(block *pbacme.Block, tx *pbacme.Transaction) bool {
	for _, event := range tx.Events {
		if p.isEventMatched(block, event) {
			return true
		}
	}

	return false
}

func (p *TransactionEventFilter) isEventMatched(block *pbacme.Block, event *pbacme.Event) bool {
	if matchedAddress, ok := p.eventFiltersByAddressAndTopic[Address(event.FromAddr)]; ok {
		// This is a keyless non-standard event, so it doesn't match with any filter
		if len(event.Keys) == 0 {
			return false
		}

		eventTopic := event.Keys[0]

		if matchedTopic, ok := matchedAddress[Topic(eventTopic)]; ok {
			for _, blockRange := range matchedTopic {
				if blockRange.endBlock == 0 {
					if block.Height >= blockRange.startBlock {
						return true
					}
				} else {
					if block.Height >= blockRange.startBlock && block.Height <= blockRange.endBlock {
						return true
					}
				}
			}
		}
	}

	return false
}
