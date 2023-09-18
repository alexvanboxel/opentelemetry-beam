package googlecloudpubsubexporter

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"
	"go.opentelemetry.io/collector/confmap"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
	"strings"
)

func Register() {
	registry.RegisterExporter("googlecloudpubsub", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	defaultConfig := googlecloudpubsubexporter.NewFactory().CreateDefaultConfig()
	err := cfg.Unmarshal(&defaultConfig)
	if err != nil {

	}
	config := defaultConfig.(*googlecloudpubsubexporter.Config)
	s = fn.CreateExporterScope(s, "googlecloudpubsub", signal, name)

	var out beam.PCollection
	switch signal {
	case "metrics":
		out = beam.ParDo(s, &PubSubMetricExporterFn{}, in)
	case "logs":
		out = beam.ParDo(s, &PubSubLogExporterFn{}, in)
	case "traces":
		out = beam.ParDo(s, &PubSubTraceExporterFn{}, in)
	}

	project, topic, err := parseTopic(config.Topic)
	if err != nil {
		return beam.PCollection{}
	}
	pubsubio.Write(s, project, topic, out)

	return beam.PCollection{}
}

func parseTopic(topic string) (string, string, error) {
	split := strings.Split(topic, "/")
	return split[1], split[3], nil
}
