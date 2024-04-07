package main

import (
	"github.com/streamingfast/firehose-acme/codec"
	pbacme "github.com/streamingfast/firehose-acme/pb/zklend/starknet/type/v1"
	firecore "github.com/streamingfast/firehose-core"
	fhCmd "github.com/streamingfast/firehose-core/cmd"
)

func main() {
	fhCmd.Main(&firecore.Chain[*pbacme.Block]{
		ShortName:            "stark",
		LongName:             "Starknet",
		ExecutableName:       "dummy-blockchain",
		FullyQualifiedModule: "github.com/streamingfast/firehose-acme",
		Version:              version,

		FirstStreamableBlock: 0,

		BlockFactory:         func() firecore.Block { return new(pbacme.Block) },
		ConsoleReaderFactory: codec.NewConsoleReader,

		Tools: &firecore.ToolsConfig[*pbacme.Block]{},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
