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

var BlockHeaderOnlyMessageName = proto.MessageName(&pbtransform.BlockHeaderOnly{})

func NewBlockHeaderOnlyTransformFactory(_ dstore.Store, _ []uint64) (*transform.Factory, error) {
	return BlockHeaderOnlyTransformFactory, nil
}

var BlockHeaderOnlyTransformFactory = &transform.Factory{
	Obj: &pbtransform.BlockHeaderOnly{},
	NewFunc: func(message *anypb.Any) (transform.Transform, error) {
		mname := message.MessageName()
		if mname != BlockHeaderOnlyMessageName {
			return nil, fmt.Errorf("expected type url %q, recevied %q ", BlockHeaderOnlyMessageName, message.TypeUrl)
		}

		filter := &pbtransform.BlockHeaderOnly{}
		err := proto.Unmarshal(message.Value, filter)
		if err != nil {
			return nil, fmt.Errorf("unexpected unmarshall error: %w", err)
		}
		return &BlockHeaderOnlyFilter{}, nil
	},
}

type BlockHeaderOnlyFilter struct{}

func (p *BlockHeaderOnlyFilter) String() string {
	return "light block filter"
}

func (p *BlockHeaderOnlyFilter) Transform(readOnlyBlk *pbbstream.Block, in transform.Input) (transform.Output, error) {
	fullBlock := &pbacme.Block{}
	err := readOnlyBlk.Payload.UnmarshalTo(fullBlock)
	if err != nil {
		return nil, fmt.Errorf("mashalling block: %w", err)
	}

	return &pbacme.Block{
		Height:    fullBlock.Height,
		Hash:      fullBlock.Hash,
		PrevHash:  fullBlock.PrevHash,
		Timestamp: fullBlock.Timestamp,
	}, nil
}
