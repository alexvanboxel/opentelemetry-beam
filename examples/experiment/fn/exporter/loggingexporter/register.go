package loggingexporter

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
)

func Register() {
	registry.RegisterExporter("logging", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, componentName, name string, cfg *confmap.Conf) beam.PCollection {
	defaultConfig := loggingexporter.NewFactory().CreateDefaultConfig()
	err := cfg.Unmarshal(&defaultConfig)
	if err != nil {
		return beam.PCollection{}
	}
	_ = defaultConfig.(*loggingexporter.Config)
	s = fn.CreateExporterScope(s, "logging", signal, name)
	switch signal {
	case "metrics":
		beam.ParDo0(s, &LoggingMetricExporterFn{}, in)
	case "logs":
		beam.ParDo0(s, &LoggingLogExporterFn{}, in)
	case "traces":
		beam.ParDo0(s, &LoggingTraceExporterFn{}, in)
	}
	return beam.PCollection{}
}
