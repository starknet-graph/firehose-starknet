package main

import (
	"fmt"
	"io"

	"github.com/streamingfast/bstream"
	pbacme "github.com/streamingfast/firehose-acme/pb/zklend/starknet/type/v1"
)

func printBlock(blk *bstream.Block, alsoPrintTransactions bool, out io.Writer) error {
	block := blk.ToProtocol().(*pbacme.Block)

	if _, err := fmt.Fprintf(out, "Block #%d (%s) (prev: %s): %d transactions\n",
		block.Height,
		block.Hash,
		block.PrevHash[0:7],
		len(block.Transactions),
	); err != nil {
		return err
	}

	if alsoPrintTransactions {
		for _, t := range block.Transactions {
			if _, err := fmt.Fprintf(out, "- Transaction %s\n", t.Hash); err != nil {
				return err
			}
		}
	}

	return nil
}
