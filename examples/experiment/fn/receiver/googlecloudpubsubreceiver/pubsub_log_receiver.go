package googlecloudpubsubreceiver

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/plog"
)

//pb "google.golang.org/genproto/googleapis/pubsub/v1"
//"go.opentelemetry.io/collector/pdata/plog"

func init() {
	register.DoFn2x0[*pb.PubsubMessage, func(logs plog.Logs)](&PubSubLogReceiverFn{})
	register.DoFn2x0[[]byte, func(logs plog.Logs)](&PubSubBytesReceiverFn{})
}

type PubSubLogReceiverFn struct {
	logsUnmarshaler plog.Unmarshaler
}

type PubSubBytesReceiverFn struct {
}

func (receiver *PubSubBytesReceiverFn) ProcessElement(elm []byte, emit func(plog.Logs)) {
}

func (receiver *PubSubLogReceiverFn) handleMetric(ctx context.Context, payload []byte, compression compression) (plog.Logs, error) {
	payload, err := decompress(payload, compression)
	if err != nil {
		return plog.Logs{}, err
	}
	return receiver.logsUnmarshaler.UnmarshalLogs(payload)
}
func (receiver *PubSubLogReceiverFn) init() {
	receiver.logsUnmarshaler = &plog.ProtoUnmarshaler{}
}

func (receiver *PubSubLogReceiverFn) ProcessElement(elm *pb.PubsubMessage, emit func(plog.Logs)) {
	if elm == nil {
		print("nothing?")
		return
	}
	if elm.Data == nil {
		print("no data?")
		return
	}
	print("at least an attribute? " + elm.Attributes["ce-id"])
	ctx := context.TODO()
	metric, err := receiver.handleMetric(ctx, elm.Data, uncompressed)
	if err != nil {
		return
	}
	emit(metric)
}

func (receiver *PubSubLogReceiverFn) Setup() {
	fmt.Println("Setup")
	receiver.logsUnmarshaler = &plog.ProtoUnmarshaler{}
}
