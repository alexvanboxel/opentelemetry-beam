package cpipe

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"go.opentelemetry.io/collector/confmap"
	"gopkg.in/yaml.v3"
	"opentelemetry-beam/examples/experiment/fn/exporter/googlecloudpubsubexporter"
	"opentelemetry-beam/examples/experiment/fn/exporter/loggingexporter"
	"opentelemetry-beam/examples/experiment/fn/processor/resourceprocessor"
	"opentelemetry-beam/examples/experiment/fn/processor/schemaprocessor"
	"opentelemetry-beam/examples/experiment/fn/receiver/googlecloudpubsubreceiver"
	"opentelemetry-beam/examples/experiment/fn/registry"
	"strings"
)

func init() {
	googlecloudpubsubreceiver.Register()
	googlecloudpubsubexporter.Register()
	loggingexporter.Register()
	schemaprocessor.Register()
	resourceprocessor.Register()
}

var (
	declaredReceivers  = map[string]*confmap.Conf{}
	declaredExporters  = map[string]*confmap.Conf{}
	declaredProcessors = map[string]*confmap.Conf{}
)

func Create(scope beam.Scope, fullConfig []byte) {
	var data map[string]any
	err := yaml.Unmarshal(fullConfig, &data)
	if err != nil {
		fmt.Print(err)
	}

	if data["service"] != nil {
		loadReceivers(data["receivers"])
		loadExporters(data["exporters"])
		loadProcessors(data["processors"])

		createPipeline(scope, data["service"].(map[string]any)["pipelines"])
	} else {
		fmt.Print("xyz")
	}

	//conf := confmap.NewFromStringMap(data)
	//component.UnmarshalConfig()

	//fmt.Print(conf)

}

func loadReceivers(ra any) {
	if ra != nil {
		r := ra.(map[string]any)
		for name, c := range r {
			conf := confmap.NewFromStringMap(c.(map[string]any))
			declaredReceivers[name] = conf
		}
	}
}

func loadExporters(ra any) {
	if ra != nil {
		r := ra.(map[string]any)
		for name, c := range r {
			conf := confmap.NewFromStringMap(c.(map[string]any))
			declaredExporters[name] = conf
		}
	}
}

func loadProcessors(ra any) {
	if ra != nil {
		r := ra.(map[string]any)
		for name, c := range r {
			conf := confmap.NewFromStringMap(c.(map[string]any))
			declaredProcessors[name] = conf
		}
	}
}

func createPipeline(scope beam.Scope, pa any) {
	if pa != nil {
		p := pa.(map[string]any)
		for signal, pipelineComponentRaw := range p {
			pipelineComponents := pipelineComponentRaw.(map[string]any)
			receivers := pipelineComponents["receivers"]
			processors := pipelineComponents["processors"]
			exporters := pipelineComponents["exporters"]

			createTransformation(scope, signal, receivers, processors, exporters)
			fmt.Println()
			_ = receivers
			_ = processors
			_ = exporters
		}

	}
}

func createComponentTransform(componentName string, declaredComponent map[string]*confmap.Conf, scope beam.Scope, pcollection beam.PCollection, signal, componentType string) beam.PCollection {
	conf := declaredComponent[componentName]
	split := strings.Split(componentName, "/")
	name := ""
	if len(split) > 1 {
		name = split[1]
	}
	factory := registry.GetFactory(split[0], componentType)
	if factory == nil {
		log.Fatalf(context.Background(), "Unable to find factory for %s, type %s", split[0], componentType)
	}
	return factory.CreateTransformFunc(scope, pcollection, signal, name, conf)
}

func createTransformation(scope beam.Scope, signal string, r any, p any, e any) {
	pcollection := beam.PCollection{}
	if r != nil {
		receivers := r.([]any)
		for _, receiverName := range receivers {
			pcollection = createComponentTransform(receiverName.(string), declaredReceivers, scope, pcollection, signal, "receivers")
		}
	}
	if p != nil {
		processors := p.([]any)
		for _, processorName := range processors {
			pcollection = createComponentTransform(processorName.(string), declaredProcessors, scope, pcollection, signal, "processors")
		}
	}
	if e != nil {
		exporter := e.([]any)
		for _, exporterName := range exporter {
			createComponentTransform(exporterName.(string), declaredExporters, scope, pcollection, signal, "exporters")
		}
	}
}
