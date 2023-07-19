package schemaprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"go.opentelemetry.io/collector/pdata/plog"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[plog.Logs, func(logs plog.Logs)](&PLogSchema{})
}

type PLogSchema struct {
	processor.LBase
}

func (fn *PLogSchema) Setup() {
	fn.Wrap(schemaprocessor.NewFactory())
}
