package otlp

import (
	"context"
	"go.opentelemetry.io/collector/pdata/plog"
)

type ProcessorLogFn struct {
}

func (fn *ProcessorLogFn) ProcessElement(ctx context.Context, pdata plog.Logs, emit func(plog.Logs)) {
	//fmt.Println(pdata.ResourceLogs())
	//log.Infof(ctx, "%s: %v", fn.LogPrefix, elm)
}
