package googlecloudpubsubreceiver

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"
	"go.opentelemetry.io/collector/confmap"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
)

func Register() {
	registry.RegisterReceiver("googlecloudpubsub", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	switch signal {
	case "metrics":
		defaultConfig := googlecloudpubsubreceiver.NewFactory().CreateDefaultConfig()
		err := cfg.Unmarshal(defaultConfig)
		if err != nil {

		}
		config := defaultConfig.(*googlecloudpubsubreceiver.Config)
		s = fn.CreateNamedScope(s, "pubsub-receiver", signal, name)
		resourceNotifications := pubsubio.Read(s, config.ProjectID, config.Subscription, &pubsubio.ReadOptions{
			Subscription:   "test-sub",
			WithAttributes: true,
		})
		return beam.ParDo(s, &PubSubReceiverFn{}, resourceNotifications)
	case "logs":
	case "traces":
	}
	log.Fatalf(context.Background(), "")
	return beam.PCollection{}
}
