package genericprocessor

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/processor"
	"gopkg.in/yaml.v3"
	"opentelemetry-beam/examples/experiment/fn"
)

func parseConfigFragment(buffer []byte, factory processor.Factory) (component.Config, error) {
	config := factory.CreateDefaultConfig()
	err := yaml.Unmarshal(buffer, config)
	if err != nil {
		log.Fatal(context.Background(), err)
		return config, err
	}
	return config, nil
}

func CreateTransform(s beam.Scope, in beam.PCollection, signal, componentName, name string, cfgFragment *confmap.Conf) beam.PCollection {
	factory := contribProcessors(componentName)

	config := factory.CreateDefaultConfig()
	err := cfgFragment.Unmarshal(&config)
	if err != nil {
		log.Fatal(context.Background(), err)
	}
	configBuffer, err := yaml.Marshal(config)
	if err != nil {
		return beam.PCollection{}
	}
	_, err = parseConfigFragment(configBuffer, factory)
	if err != nil {
		return beam.PCollection{}
	}

	s = fn.CreateProcessorScope(s, componentName, signal, name)
	switch signal {
	case "metrics":
		return beam.ParDo(s, &GenericMetricProcessor{
			Base: Base{
				Name:   componentName,
				Config: configBuffer,
			},
		}, in)
	case "logs":
		return beam.ParDo(s, &GenericLogProcessor{
			Base: Base{
				Name:   componentName,
				Config: configBuffer,
			},
		}, in)
	case "traces":
		return beam.ParDo(s, &GenericTraceProcessor{
			Base: Base{
				Name:   componentName,
				Config: configBuffer,
			},
		}, in)
	}
	log.Fatalf(context.Background(), "unknown signal %s", signal)
	return beam.PCollection{}
}
