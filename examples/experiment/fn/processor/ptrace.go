package processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

type TBase struct {
	Base
	emit            func(ptrace.Traces)
	tracesProcessor processor.Traces
}

func (fn *TBase) ProcessElement(element ptrace.Traces, emit func(ptrace.Traces)) {
	fn.emit = emit
	err := fn.tracesProcessor.ConsumeTraces(context.Background(), element)
	if err != nil {
		return
	}
}

func (fn *TBase) ConsumeTraces(ctx context.Context, m ptrace.Traces) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *TBase) Teardown() {
	fn.tracesProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *TBase) Wrap(factory processor.Factory) {
	fmt.Println("Setup")

	tracesProcessor, err := factory.CreateTracesProcessor(context.Background(), fn.createSettings(), fn.Base.parseStringConfig(factory.CreateDefaultConfig()), fn)
	if err != nil {
		return
	}
	fn.tracesProcessor = tracesProcessor
	tracesProcessor.Start(context.Background(), &BeamHost{})
}
