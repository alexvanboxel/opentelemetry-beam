package schemaprocessor

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"go.opentelemetry.io/collector/confmap"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/processor"
	"opentelemetry-beam/examples/experiment/fn/registry"
)

func Register() {
	registry.RegisterProcessor("schema", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	s = fn.CreateProcessorScope(s, "schema", signal, name)
	switch signal {
	case "metrics":
		return beam.ParDo(s, &PMetricSchema{
			processor.MBase{
				Base: processor.Base{
					Config: cfg,
				},
			},
		}, in)
	case "logs":
		return beam.ParDo(s, &PLogSchema{
			processor.LBase{
				Base: processor.Base{
					Config: cfg,
				},
			},
		}, in)
	case "traces":
		return beam.ParDo(s, &PTraceSchema{
			processor.TBase{
				Base: processor.Base{
					Config: cfg,
				},
			},
		}, in)
	}
	log.Fatalf(context.Background(), "unknown signal %s", signal)
	return beam.PCollection{}
}
