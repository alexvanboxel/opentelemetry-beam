package resourceprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"go.opentelemetry.io/collector/pdata/plog"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[plog.Logs, func(logs plog.Logs)](&PLogResource{})
}

type PLogResource struct {
	processor.LBase
}

func (fn *PLogResource) Setup() {
	fn.Wrap(resourceprocessor.NewFactory())
}
