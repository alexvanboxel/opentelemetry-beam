package genericprocessor

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"log"
)

func init() {
	register.DoFn2x0[ptrace.Traces, func(ptrace.Traces)](&GenericTraceProcessor{})
}

type GenericTraceProcessor struct {
	Base
	emit            func(ptrace.Traces)
	tracesProcessor processor.Traces
}

func (fn *GenericTraceProcessor) ProcessElement(element ptrace.Traces, emit func(ptrace.Traces)) {
	fn.emit = emit
	err := fn.tracesProcessor.ConsumeTraces(context.Background(), element)
	if err != nil {
		return
	}
}

func (fn *GenericTraceProcessor) ConsumeTraces(ctx context.Context, m ptrace.Traces) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *GenericTraceProcessor) Teardown() {
	fn.tracesProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *GenericTraceProcessor) Setup() {
	factory := contribProcessors(fn.Name)
	config, err := parseConfigFragment(fn.Config, factory)
	if err != nil {
		log.Fatal(context.Background(), err)
	}

	tracesProcessor, err := factory.CreateTracesProcessor(context.Background(), fn.createSettings(), config, fn)
	if err != nil {
		return
	}
	fn.tracesProcessor = tracesProcessor
	tracesProcessor.Start(context.Background(), &BeamHost{})
}
