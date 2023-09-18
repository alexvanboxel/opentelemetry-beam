module opentelemetry-beam

go 1.18

require (
	cloud.google.com/go/pubsub v1.33.0
	github.com/apache/beam/sdks/v2 v2.50.0
	github.com/golang/protobuf v1.5.3 // TODO(danoliveira): Fully replace this with google.golang.org/protobuf
	google.golang.org/protobuf v1.31.0
)
