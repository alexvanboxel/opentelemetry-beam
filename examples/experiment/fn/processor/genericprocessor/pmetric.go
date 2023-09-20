package genericprocessor

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"
	"log"
)

func init() {
	register.DoFn2x0[pmetric.Metrics, func(pmetric.Metrics)](&GenericMetricProcessor{})
}

type GenericMetricProcessor struct {
	Base
	emit             func(pmetric.Metrics)
	metricsProcessor processor.Metrics
}

func (fn *GenericMetricProcessor) ProcessElement(element pmetric.Metrics, emit func(pmetric.Metrics)) {
	fn.emit = emit
	err := fn.metricsProcessor.ConsumeMetrics(context.Background(), element)
	if err != nil {
		return
	}
}

// ConsumeMetrics is an implementation of
func (fn *GenericMetricProcessor) ConsumeMetrics(ctx context.Context, m pmetric.Metrics) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *GenericMetricProcessor) Teardown() {
	fn.metricsProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *GenericMetricProcessor) Setup() {
	factory := contribProcessors(fn.Name)
	config, err := parseConfigFragment(fn.Config, factory)
	if err != nil {
		log.Fatal(context.Background(), err)
	}

	metricsProcessor, err := factory.CreateMetricsProcessor(context.Background(), fn.createSettings(), config, fn)
	if err != nil {
		return
	}
	fn.metricsProcessor = metricsProcessor
	metricsProcessor.Start(context.Background(), &BeamHost{})
}
