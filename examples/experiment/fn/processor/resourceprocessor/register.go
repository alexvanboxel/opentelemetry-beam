package resourceprocessor

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
	registry.RegisterProcessor("resource", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	s = fn.CreateNamedScope(s, "resource", signal, name)
	switch signal {
	case "metrics":
		return beam.ParDo(s, &PMetricResource{
			processor.MBase{
				Base: processor.Base{
					Config: cfg,
				},
			},
		}, in)
	case "logs":
		return beam.ParDo(s, &PLogResource{
			processor.LBase{
				Base: processor.Base{
					Config: cfg,
				},
			},
		}, in)
	case "traces":
		return beam.ParDo(s, &PTraceResource{
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
