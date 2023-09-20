package genericprocessor

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	"go.opentelemetry.io/collector/processor"
)

func contribProcessors(name string) processor.Factory {
	switch name {
	case "attributes":
		return attributesprocessor.NewFactory()
	case "filter":
		return filterprocessor.NewFactory()
	case "resource":
		return resourceprocessor.NewFactory()
	case "schema":
		return schemaprocessor.NewFactory()
	case "transform":
		return transformprocessor.NewFactory()

	}
	return coreProcessors(name)
}

func coreProcessors(name string) processor.Factory {
	return nil
}
