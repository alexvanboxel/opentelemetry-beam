package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"log"
	"opentelemetry-beam/examples/experiment/cpipe"
	"os"
	"path/filepath"
)

var (
	// Optional flag with the pubsub topic of your FHIR store to read and log upon store updates.
	pubsubTopic = flag.String("pubsubTopic", "otlp-ingress-metrics", "The PubSub topic to listen to.")
)

func main() {
	beam.Init()
	flag.Parse()
	_ = *gcpopts.Project
	pipeline, scope := beam.NewPipelineWithRoot()

	fileName := "/Users/alex.vanboxel/src/open-telemetry/opentelemetry-collector-collibra/dist/config/beam-experiment.yaml"
	content, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		fmt.Print(err)
	}

	cpipe.Create(scope, content)

	_ = pipeline
	ctx := context.Background()
	if err := beamx.Run(ctx, pipeline); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

func CreatePipeline() {

}

func fixed() {
	//	flag.Parse()
	//	beam.Init()
	//
	//	p, s := beam.NewPipelineWithRoot()
	//
	//	resourceNotifications := pubsubio.Read(s, *gcpopts.Project, *pubsubTopic, &pubsubio.ReadOptions{
	//		Subscription:   "test-sub",
	//		WithAttributes: true,
	//	})
	//	metrics := beam.ParDo(s, &receiver.PubSubReceiverFn{}, resourceNotifications)
	//	out := beam.ParDo(s, pwrapper.NewMetricResourceWrapper(`
	//attributes:
	//- key: deployment.environment
	//  value: beam
	//  action: upsert`), metrics)
	//	_ = beam.ParDo(s, &exporter.PubSubExporterFn{}, out)
	//
	//	ctx := context.Background()
	//	if err := beamx.Run(ctx, p); err != nil {
	//		log.Fatalf(ctx, "Failed to execute job: %v", err)
	//	}
}
