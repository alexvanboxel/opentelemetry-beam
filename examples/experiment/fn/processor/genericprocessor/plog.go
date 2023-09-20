package genericprocessor

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"log"
)

func init() {
	register.DoFn2x0[plog.Logs, func(logs plog.Logs)](&GenericLogProcessor{})
}

type GenericLogProcessor struct {
	Base
	emit          func(plog.Logs)
	logsProcessor processor.Logs
}

func (fn *GenericLogProcessor) ProcessElement(element plog.Logs, emit func(plog.Logs)) {
	fn.emit = emit
	err := fn.logsProcessor.ConsumeLogs(context.Background(), element)
	if err != nil {
		return
	}
}

func (fn *GenericLogProcessor) ConsumeLogs(ctx context.Context, m plog.Logs) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *GenericLogProcessor) Teardown() {
	fn.logsProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *GenericLogProcessor) Setup() {
	factory := contribProcessors(fn.Name)
	config, err := parseConfigFragment(fn.Config, factory)
	if err != nil {
		log.Fatal(context.Background(), err)
	}

	logsProcessor, err := factory.CreateLogsProcessor(context.Background(), fn.createSettings(), config, fn)
	if err != nil {
		return
	}
	fn.logsProcessor = logsProcessor
	logsProcessor.Start(context.Background(), &BeamHost{})
}
