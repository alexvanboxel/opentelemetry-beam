package loggingexporter

import (
	"bytes"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"compress/gzip"
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
	register.DoFn2x0[pmetric.Metrics, func(*pb.PubsubMessage)](&PubSubExporterFn{})
}

type PubSubExporterFn struct {
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

func (receiver *PubSubExporterFn) ProcessElement(metrics pmetric.Metrics, emit func(elm *pb.PubsubMessage)) {
	rls := metrics.ResourceMetrics()
	for i := 0; i < rls.Len(); i++ {
		rs := rls.At(i)
		for k, v := range rs.Resource().Attributes().AsRaw() {
			fmt.Printf("ResourceAttr: %s=%v\n", k, v)
		}
		ilss := rs.ScopeMetrics()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			logs := ils.Metrics()
			for k := 0; k < logs.Len(); k++ {
				metric := logs.At(k)
				fmt.Printf("MetricName: %s (%s)\n", metric.Name(), metric.Description())
			}
		}
	}

}

func (receiver *PubSubExporterFn) Setup() {
	fmt.Println("Setup")
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}
