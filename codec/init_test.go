package codec

import (
	"github.com/streamingfast/bstream"
	"github.com/streamingfast/logging"
)

var zlog, _ = logging.PackageLogger("codec", "github.com/streamingfast/firehose-acme/codec_test")

func init() {
	logging.InstantiateLoggers()

	bstream.GetBlockPayloadSetter = bstream.MemoryBlockPayloadSetter
}
