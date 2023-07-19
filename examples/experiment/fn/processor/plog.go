package processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
)

type LBase struct {
	Base
	emit          func(plog.Logs)
	logsProcessor processor.Logs
}

func (fn *LBase) ProcessElement(element plog.Logs, emit func(plog.Logs)) {
	fn.emit = emit
	err := fn.logsProcessor.ConsumeLogs(context.Background(), element)
	if err != nil {
		return
	}
}

func (fn *LBase) ConsumeLogs(ctx context.Context, m plog.Logs) error {
	fmt.Println("emit(m)")
	fn.emit(m)
	return nil
}

func (fn *LBase) Teardown() {
	fn.logsProcessor.Shutdown(context.Background())
	fmt.Println("Teardown")
}

func (fn *LBase) Wrap(factory processor.Factory) {
	fmt.Println("Setup")

	logsProcessor, err := factory.CreateLogsProcessor(context.Background(), fn.createSettings(), fn.Base.parseStringConfig(factory.CreateDefaultConfig()), fn)
	if err != nil {
		return
	}
	fn.logsProcessor = logsProcessor
	logsProcessor.Start(context.Background(), &BeamHost{})
}
