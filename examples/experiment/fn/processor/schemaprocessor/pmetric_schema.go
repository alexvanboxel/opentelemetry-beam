package schemaprocessor

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"opentelemetry-beam/examples/experiment/fn/processor"
)

func init() {
	register.DoFn2x0[pmetric.Metrics, func(pmetric.Metrics)](&PMetricSchema{})
}

type PMetricSchema struct {
	processor.MBase
}

func (fn *PMetricSchema) Setup() {
	fn.Wrap(schemaprocessor.NewFactory())
}
