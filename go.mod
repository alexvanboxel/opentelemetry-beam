module opentelemetry-beam

go 1.18

require (
	github.com/apache/beam/sdks/v2 v2.47.0
	cloud.google.com/go/pubsub v1.24.0
	github.com/golang/protobuf v1.5.2 // TODO(danoliveira): Fully replace this with google.golang.org/protobuf
	google.golang.org/protobuf v1.28.1
)
