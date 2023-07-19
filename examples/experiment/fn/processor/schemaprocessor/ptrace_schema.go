package schemaprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[ptrace.Traces, func(ptrace.Traces)](&PTraceSchema{})
}

type PTraceSchema struct {
	processor.TBase
}

func (fn *PTraceSchema) Setup() {
	fn.Wrap(schemaprocessor.NewFactory())
}
