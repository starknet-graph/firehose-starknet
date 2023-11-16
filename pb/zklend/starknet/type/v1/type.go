package pbacme

import (
	"fmt"
	"time"
)

func (b *Block) GetFirehoseBlockID() string {
	return fmt.Sprintf("0x%x", b.Hash)
}

func (b *Block) GetFirehoseBlockNumber() uint64 {
	return b.Height
}

func (b *Block) GetFirehoseBlockParentNumber() uint64 {
	return b.Height - 1
}

func (b *Block) GetFirehoseBlockParentID() string {
	return fmt.Sprintf("0x%x", b.PrevHash)
}

func (b *Block) GetFirehoseBlockTime() time.Time {
	return time.Unix(int64(b.Timestamp), 0).UTC()
}

func (b *Block) GetFirehoseBlockVersion() int32 {
	return 0
}

func (b *Block) GetFirehoseBlockLIBNum() uint64 {
	if b.Height <= 100 {
		return 0
	} else {
		return b.Height - 100
	}
}
