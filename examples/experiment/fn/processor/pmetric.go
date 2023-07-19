package processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"
)

type MBase struct {
	Base
	emit             func(pmetric.Metrics)
	metricsProcessor processor.Metrics
}

func (fn *MBase) ProcessElement(element pmetric.Metrics, emit func(pmetric.Metrics)) {
	fn.emit = emit
	err := fn.metricsProcessor.ConsumeMetrics(context.Background(), element)
	if err != nil {
		return
	}
}

func (fn *MBase) ConsumeMetrics(ctx context.Context, m pmetric.Metrics) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *MBase) Teardown() {
	fn.metricsProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *MBase) Wrap(factory processor.Factory) {
	fmt.Println("Setup")

	metricsProcessor, err := factory.CreateMetricsProcessor(context.Background(), fn.createSettings(), fn.Base.parseStringConfig(factory.CreateDefaultConfig()), fn)
	if err != nil {
		return
	}
	fn.metricsProcessor = metricsProcessor
	metricsProcessor.Start(context.Background(), &BeamHost{})
}
