package types

import (
	"fmt"

	pbacme "github.com/starknet-graph/firehose-starknet/types/pb/zklend/starknet/type/v1"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
	"google.golang.org/protobuf/proto"
)

func BlockFromProto(b *pbacme.Block) (*bstream.Block, error) {
	content, err := proto.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal to binary form: %s", err)
	}

	blockNumber := b.Number()

	block := &bstream.Block{
		Id:             b.ID(),
		Number:         blockNumber,
		PreviousId:     b.PreviousID(),
		Timestamp:      b.Time(),
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}

	// For simpliciy's sake we're pretending StarkNet cannot re-org for more than 10 blocks
	if blockNumber <= 10 {
		block.LibNum = 0
	} else {

		block.LibNum = blockNumber - 10
	}

	return bstream.GetBlockPayloadSetter(block, content)
}
