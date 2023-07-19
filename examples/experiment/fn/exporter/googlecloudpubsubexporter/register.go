package googlecloudpubsubexporter

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"
	"go.opentelemetry.io/collector/confmap"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
)

func Register() {
	registry.RegisterExporter("googlecloudpubsub", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	defaultConfig := googlecloudpubsubexporter.NewFactory().CreateDefaultConfig()
	err := cfg.Unmarshal(&defaultConfig)
	if err != nil {

	}
	config := defaultConfig.(googlecloudpubsubexporter.Config)
	_ = config
	switch signal {
	case "metrics":
		s = fn.CreateNamedScope(s, "pubsub-exporter", signal, name)
		return beam.ParDo(s, &PubSubExporterFn{}, in)
	case "logs":
	case "traces":
	}
	log.Fatalf(context.Background(), "")
	return beam.PCollection{}
}
