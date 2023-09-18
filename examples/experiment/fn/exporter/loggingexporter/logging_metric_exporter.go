package loggingexporter

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func init() {
	register.DoFn1x0[pmetric.Metrics](&LoggingMetricExporterFn{})
}

type LoggingMetricExporterFn struct {
	metricsUnmarshaler pmetric.Unmarshaler
}

func (receiver *LoggingMetricExporterFn) ProcessElement(metrics pmetric.Metrics) {
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

func (receiver *LoggingMetricExporterFn) Setup() {
	fmt.Println("Setup")
	receiver.metricsUnmarshaler = &pmetric.ProtoUnmarshaler{}
}
