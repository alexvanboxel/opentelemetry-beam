package loggingexporter

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
)

func Register() {
	registry.RegisterExporter("logging", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	defaultConfig := loggingexporter.NewFactory().CreateDefaultConfig()
	err := cfg.Unmarshal(&defaultConfig)
	if err != nil {

	}
	config := defaultConfig.(*loggingexporter.Config)
	_ = config
	switch signal {
	case "metrics":
		s = fn.CreateNamedScope(s, "logging", signal, name)
		return beam.ParDo(s, &PubSubExporterFn{}, in)
	case "logs":
	case "traces":
	}
	log.Fatalf(context.Background(), "")
	return beam.PCollection{}
}
