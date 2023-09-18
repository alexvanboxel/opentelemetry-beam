package loggingexporter

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func init() {
	register.DoFn1x0[ptrace.Traces](&LoggingTraceExporterFn{})
}

type LoggingTraceExporterFn struct {
}

func (receiver *LoggingTraceExporterFn) ProcessElement(traces ptrace.Traces) {
	rls := traces.ResourceSpans()
	for i := 0; i < rls.Len(); i++ {
		rs := rls.At(i)
		for k, v := range rs.Resource().Attributes().AsRaw() {
			fmt.Printf("ResourceAttr: %s=%v\n", k, v)
		}
		ilss := rs.ScopeSpans()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			spans := ils.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				fmt.Printf("SpanName: %s\n", span.Name())
			}
		}
	}
}

func (receiver *LoggingTraceExporterFn) Setup() {
	fmt.Println("Setup")
}
