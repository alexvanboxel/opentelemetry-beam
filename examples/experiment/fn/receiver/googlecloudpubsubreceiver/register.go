package googlecloudpubsubreceiver

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"
	"go.opentelemetry.io/collector/confmap"
	"opentelemetry-beam/examples/experiment/fn"
	"opentelemetry-beam/examples/experiment/fn/registry"
	"strings"
)

func Register() {
	registry.RegisterReceiver("googlecloudpubsub", CreateTransform)
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, name string, cfg *confmap.Conf) beam.PCollection {
	defaultConfig := googlecloudpubsubreceiver.NewFactory().CreateDefaultConfig()
	err := cfg.Unmarshal(defaultConfig)
	if err != nil {

	}
	config := defaultConfig.(*googlecloudpubsubreceiver.Config)
	project, topic, subscription, err := discoverTopicSubscription(config.ProjectID, config.Subscription)
	if err != nil {
		return beam.PCollection{}
	}
	s = fn.CreateReceiverScope(s, "googlecloudpubsub", signal, name)
	// TODO: the real receiver is able to split a single pubsub topic in it's components
	resourceNotifications := pubsubio.Read(s, project, topic, &pubsubio.ReadOptions{
		Subscription:   subscription,
		WithAttributes: true,
	})
	switch signal {
	case "metrics":
		return beam.ParDo(s, &PubSubMetricReceiverFn{}, resourceNotifications)
	case "logs":
		return beam.ParDo(s, &PubSubLogReceiverFn{}, resourceNotifications)
	case "traces":
		return beam.ParDo(s, &PubSubTraceReceiverFn{}, resourceNotifications)
	}
	log.Fatalf(context.Background(), "")
	return beam.PCollection{}
}

func discoverTopicSubscription(projectID, subscription string) (string, string, string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return "", "", "", err
	}
	subscriptionSplit := strings.Split(subscription, "/")
	s := client.SubscriptionInProject(subscriptionSplit[3], subscriptionSplit[1])
	config, err := s.Config(ctx)
	if err != nil {
		return "", "", "", err
	}

	topic := config.Topic.String()
	topicSplit := strings.Split(topic, "/")

	return subscriptionSplit[1], topicSplit[3], subscriptionSplit[3], nil
}
