package googlecloudpubsubexporter

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func init() {
	register.DoFn2x0[pmetric.Metrics, func(*pb.PubsubMessage)](&PubSubMetricExporterFn{})
}

type PubSubMetricExporterFn struct {
	marshaler pmetric.Marshaler
}

func (receiver *PubSubMetricExporterFn) ProcessElement(metrics pmetric.Metrics, emit func(*pb.PubsubMessage)) {
	buffer, err := receiver.marshaler.MarshalMetrics(metrics)
	if err != nil {
		return
	}
	message := pb.PubsubMessage{
		Data: buffer,
		Attributes: map[string]string{
			"ce-type":          "org.opentelemetry.otlp.metrics.v1",
			"content-type":     "application/protobuf",
			"content-encoding": "",
		},
	}
	emit(&message)
}

func (receiver *PubSubMetricExporterFn) Setup() {
	receiver.marshaler = &pmetric.ProtoMarshaler{}
}
