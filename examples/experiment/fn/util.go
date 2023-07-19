package fn

import (
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
)

func CreateNamedScope(scope beam.Scope, componentName, signal, name string) beam.Scope {
	if name == "" {
		return scope.Scope(fmt.Sprintf("%s.%s", componentName, signal))
	}
	return scope.Scope(fmt.Sprintf("%s.%s.%s", componentName, name, signal))
}
