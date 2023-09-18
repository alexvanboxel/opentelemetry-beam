package googlecloudpubsubexporter

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func init() {
	register.DoFn2x0[ptrace.Traces, func(*pb.PubsubMessage)](&PubSubTraceExporterFn{})
}

type PubSubTraceExporterFn struct {
	marshaler ptrace.Marshaler
}

func (receiver *PubSubTraceExporterFn) ProcessElement(metrics ptrace.Traces, emit func(*pb.PubsubMessage)) {
	buffer, err := receiver.marshaler.MarshalTraces(metrics)
	if err != nil {
		return
	}
	message := pb.PubsubMessage{
		Data: buffer,
		Attributes: map[string]string{
			"ce-type":          "org.opentelemetry.otlp.traces.v1",
			"content-type":     "application/protobuf",
			"content-encoding": "",
		},
	}
	emit(&message)
}

func (receiver *PubSubTraceExporterFn) Setup() {
	receiver.marshaler = &ptrace.ProtoMarshaler{}
}
