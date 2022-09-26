package pbacme

import (
	"fmt"
	"time"
)

func (b *Block) ID() string {
	return fmt.Sprintf("0x%x", b.Hash)
}

func (b *Block) Number() uint64 {
	return b.Height
}

func (b *Block) PreviousID() string {
	return fmt.Sprintf("0x%x", b.PrevHash)
}

func (b *Block) Time() time.Time {
	return time.Unix(int64(b.Timestamp), 0).UTC()
}
