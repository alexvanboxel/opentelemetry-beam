package loggingexporter

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func init() {
	register.DoFn1x0[plog.Logs](&LoggingLogExporterFn{})
}

type LoggingLogExporterFn struct {
	metricsUnmarshaler pmetric.Unmarshaler
}

func (receiver *LoggingLogExporterFn) ProcessElement(metrics plog.Logs) {
	rls := metrics.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rs := rls.At(i)
		for k, v := range rs.Resource().Attributes().AsRaw() {
			fmt.Printf("ResourceAttr: %s=%v\n", k, v)
		}
		ilss := rs.ScopeLogs()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			logs := ils.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				metric := logs.At(k)
				fmt.Printf("Body: %s\n", metric.Body())
			}
		}
	}
}

func (receiver *LoggingLogExporterFn) Setup() {
	fmt.Println("Setup")
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}
