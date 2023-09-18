package resourceprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[pmetric.Metrics, func(pmetric.Metrics)](&PMetricResource{})
}

type PMetricResource struct {
	processor.MBase
}

func (fn *PMetricResource) Setup() {
	fn.Wrap(resourceprocessor.NewFactory())
}
