package googlecloudpubsubreceiver

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

//pb "google.golang.org/genproto/googleapis/pubsub/v1"
//"go.opentelemetry.io/collector/pdata/plog"

func init() {
	register.DoFn2x0[*pb.PubsubMessage, func(traces ptrace.Traces)](&PubSubTraceReceiverFn{})
}

type PubSubTraceReceiverFn struct {
	tracesUnmarshaler ptrace.Unmarshaler
}

func (receiver *PubSubTraceReceiverFn) handleMetric(ctx context.Context, payload []byte, compression compression) (ptrace.Traces, error) {
	payload, err := decompress(payload, compression)
	if err != nil {
		return ptrace.Traces{}, err
	}
	return receiver.tracesUnmarshaler.UnmarshalTraces(payload)
}
func (receiver *PubSubTraceReceiverFn) init() {
	receiver.tracesUnmarshaler = &ptrace.ProtoUnmarshaler{}
}

func (receiver *PubSubTraceReceiverFn) ProcessElement(elm *pb.PubsubMessage, emit func(ptrace.Traces)) {
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

func (receiver *PubSubTraceReceiverFn) Setup() {
	fmt.Println("Setup")
	receiver.tracesUnmarshaler = &ptrace.ProtoUnmarshaler{}
}
