package googlecloudpubsubreceiver

import (
	"bytes"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"io"
)

//pb "google.golang.org/genproto/googleapis/pubsub/v1"
//"go.opentelemetry.io/collector/pdata/plog"

type compression int

const (
	uncompressed compression = iota
	gZip                     = iota
)

func init() {
	register.DoFn2x0[*pb.PubsubMessage, func(pmetric.Metrics)](&PubSubReceiverFn{})
}

type PubSubReceiverFn struct {
	metricsUnmarshaler pmetric.Unmarshaler
}

func decompress(payload []byte, compression compression) ([]byte, error) {
	if compression == gZip {
		reader, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			return nil, err
		}
		return io.ReadAll(reader)
	}
	return payload, nil
}

func (receiver *PubSubReceiverFn) handleMetric(ctx context.Context, payload []byte, compression compression) (pmetric.Metrics, error) {
	payload, err := decompress(payload, compression)
	if err != nil {
		return pmetric.Metrics{}, err
	}
	return receiver.metricsUnmarshaler.UnmarshalMetrics(payload)
}
func (receiver *PubSubReceiverFn) init() {
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}

func (receiver *PubSubReceiverFn) ProcessElement(elm *pb.PubsubMessage, emit func(pmetric.Metrics)) {
	if elm == nil {
		print("nothing?")
		return
	}
	if elm.Data == nil {
		print("no data?")
		return
	}
	print(elm.Attributes["ce-id"])
	ctx := context.TODO()
	metric, err := receiver.handleMetric(ctx, elm.Data, gZip)
	if err != nil {
		return
	}
	emit(metric)
	//log.Infof(ctx, "%s: %v", fn.LogPrefix, elm)

}

func (receiver *PubSubReceiverFn) Setup() {
	fmt.Println("Setup")
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}
