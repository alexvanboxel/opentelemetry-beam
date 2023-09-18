package fn

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
)

func CreateProcessorScope(scope beam.Scope, componentName, signal, name string) beam.Scope {
	if name == "" {
		return scope.Scope(fmt.Sprintf("%s.%s[p]", componentName, signal))
	}
	return scope.Scope(fmt.Sprintf("%s.%s.%s[p]", componentName, name, signal))
}

func CreateReceiverScope(scope beam.Scope, componentName, signal, name string) beam.Scope {
	if name == "" {
		return scope.Scope(fmt.Sprintf("%s.%s[r]", componentName, signal))
	}
	return scope.Scope(fmt.Sprintf("%s.%s.%s[r]", componentName, name, signal))
}

func CreateExporterScope(scope beam.Scope, componentName, signal, name string) beam.Scope {
	if name == "" {
		return scope.Scope(fmt.Sprintf("%s.%s[e]", componentName, signal))
	}
	return scope.Scope(fmt.Sprintf("%s.%s.%s[e]", componentName, name, signal))
}
