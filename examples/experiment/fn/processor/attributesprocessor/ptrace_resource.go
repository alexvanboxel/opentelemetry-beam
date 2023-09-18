package resourceprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[ptrace.Traces, func(ptrace.Traces)](&PTraceResource{})
}

type PTraceResource struct {
	processor.TBase
}

func (fn *PTraceResource) Setup() {
	fn.Wrap(resourceprocessor.NewFactory())
}
