package registry

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"go.opentelemetry.io/collector/confmap"
)

type CreateTransformFunc func(s beam.Scope, in beam.PCollection, signal, componentName, name string, cfg *confmap.Conf) beam.PCollection

type FactoryFn func(cfg *confmap.Conf) any

type Factory struct {
	Ignore              bool
	NotSupported        bool
	CreateTransformFunc CreateTransformFunc
}
