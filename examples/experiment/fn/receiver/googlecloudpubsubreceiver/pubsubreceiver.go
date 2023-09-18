package googlecloudpubsubreceiver

import (
	"bytes"
	"compress/gzip"
	"io"
)

//pb "google.golang.org/genproto/googleapis/pubsub/v1"
//"go.opentelemetry.io/collector/pdata/plog"

type compression int

const (
	uncompressed compression = iota
	gZip                     = iota
)

func decompress(payload []byte, compression compression) ([]byte, error) {
	switch compression {
	case gZip:
		reader, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			return nil, err
		}
		return io.ReadAll(reader)
	case uncompressed:
		return payload, nil
	}
	return payload, nil
}
