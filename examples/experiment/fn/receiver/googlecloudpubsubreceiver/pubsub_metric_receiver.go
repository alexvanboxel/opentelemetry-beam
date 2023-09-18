package googlecloudpubsubreceiver

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

//pb "google.golang.org/genproto/googleapis/pubsub/v1"
//"go.opentelemetry.io/collector/pdata/plog"

func init() {
	register.DoFn2x0[*pb.PubsubMessage, func(pmetric.Metrics)](&PubSubMetricReceiverFn{})
}

type PubSubMetricReceiverFn struct {
	metricsUnmarshaler pmetric.Unmarshaler
}

func (receiver *PubSubMetricReceiverFn) handleMetric(ctx context.Context, payload []byte, compression compression) (pmetric.Metrics, error) {
	payload, err := decompress(payload, compression)
	if err != nil {
		return pmetric.Metrics{}, err
	}
	return receiver.metricsUnmarshaler.UnmarshalMetrics(payload)
}
func (receiver *PubSubMetricReceiverFn) init() {
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}

func (receiver *PubSubMetricReceiverFn) ProcessElement(elm *pb.PubsubMessage, emit func(pmetric.Metrics)) {
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

func (receiver *PubSubMetricReceiverFn) Setup() {
	fmt.Println("Setup")
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}
