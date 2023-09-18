package googlecloudpubsubexporter

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/plog"
)

func init() {
	register.DoFn2x0[plog.Logs, func(*pb.PubsubMessage)](&PubSubLogExporterFn{})
}

type PubSubLogExporterFn struct {
	marshaler plog.Marshaler
}

func (receiver *PubSubLogExporterFn) ProcessElement(metrics plog.Logs, emit func(*pb.PubsubMessage)) {
	buffer, err := receiver.marshaler.MarshalLogs(metrics)
	if err != nil {
		return
	}
	message := pb.PubsubMessage{
		Data: buffer,
		Attributes: map[string]string{
			"ce-type":          "org.opentelemetry.otlp.logs.v1",
			"content-type":     "application/protobuf",
			"content-encoding": "",
		},
	}
	emit(&message)
}

func (receiver *PubSubLogExporterFn) Setup() {
	receiver.marshaler = &plog.ProtoMarshaler{}
}
