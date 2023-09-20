package genericprocessor

import (
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	metricnoop "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Base struct {
	Name   string
	Config []byte
}

func (fn *Base) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (fn *Base) createSettings() processor.CreateSettings {
	development, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf(context.Background(), "")
	}
	return processor.CreateSettings{
		ID: component.NewID("beam"),
		TelemetrySettings: component.TelemetrySettings{
			Logger:         development,
			TracerProvider: &BeamTraceProvider{},
			MeterProvider:  &BeamMeterProvider{},
			MetricsLevel:   0,
		},
		BuildInfo: component.BuildInfo{},
	}
}

type BeamMeterProvider struct {
	embedded.MeterProvider
}

func (b BeamMeterProvider) Meter(instrumentationName string, opts ...metric.MeterOption) metric.Meter {
	return metricnoop.NewMeterProvider().Meter(instrumentationName, opts...)
}

type BeamTraceProvider struct {
}

func (b BeamTraceProvider) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	return trace.NewNoopTracerProvider().Tracer("noop")
}

type BeamHost struct {
}

func (b BeamHost) ReportFatalError(err error) {
	//TODO implement me
	panic("implement me")
}

func (b BeamHost) GetFactory(kind component.Kind, componentType component.Type) component.Factory {
	//TODO implement me
	panic("implement me")
}

func (b BeamHost) GetExtensions() map[component.ID]component.Component {
	//TODO implement me
	panic("implement me")
}

func (b BeamHost) GetExporters() map[component.DataType]map[component.ID]component.Component {
	//TODO implement me
	panic("implement me")
}
