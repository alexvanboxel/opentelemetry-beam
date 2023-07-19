package registry

var receivers = map[string]*Factory{}
var processors = map[string]*Factory{}
var exporters = map[string]*Factory{}

func Register() {
	processors = map[string]*Factory{
		// https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor
		"batch": &Factory{
			Ignore: true, // TODO, native implementation
		},
		"memorylimiter": &Factory{
			Ignore: true, // memory management done by beam
		},
		"cumulativetodelta": &Factory{
			NotSupported: true, // TODO, native implementation
		},
		"datadog": &Factory{
			NotSupported: true,
		},
		"deltatorate": &Factory{
			NotSupported: true, // TODO, native implementation
		},
		"k8sattributes": &Factory{
			NotSupported: true, // needs to run in a cluster
		},
		"resourcedetection": &Factory{
			NotSupported: true, // can't detect the resource remotely
		},
		"routing": &Factory{
			NotSupported: true,
		}, // TODO but tricky
		"servicegraph": &Factory{
			NotSupported: true,
		},
		"spanmetrics": &Factory{
			NotSupported: true, // TODO, native implementation
		},
		"tailsampling": &Factory{
			NotSupported: true, // TODO, native implementation
		},
		"transform": &Factory{
			//FactoryFn: transformprocessor.NewFactory(),
		}, // TODO
	}
	// TODO wrap
	/*
		"probabilisticsampler": &Factory{}, // TODO
			"redaction":            &Factory{}, // TODO
				"groupbyattrs": &Factory{}, // TODO
				"groupbytrace": &Factory{}, // TODO
				"filter": &Factory{
				"attributes": &Factory{
					//FactoryFn: attributesFactory.NewFactory(),
				},
				"logstransform":        &Factory{},
				"metricsgeneration":    &Factory{},
				"metricstransform":     &Factory{},
				"span": &Factory{
					FactoryFn: spanFactory.NewFactory(),
				},
	*/
}

func RegisterProcessor(name string, factoryFn CreateTransformFunc) {
	_, found := processors[name]
	if !found {
		processors[name] = &Factory{}
	}
	processors[name].CreateTransformFunc = factoryFn
}

func RegisterExporter(name string, factoryFn CreateTransformFunc) {
	_, found := exporters[name]
	if !found {
		exporters[name] = &Factory{}
	}
	exporters[name].CreateTransformFunc = factoryFn
}

func RegisterReceiver(name string, factoryFn CreateTransformFunc) {
	_, found := receivers[name]
	if !found {
		receivers[name] = &Factory{}
	}
	receivers[name].CreateTransformFunc = factoryFn
}

func GetFactory(name string, componentType string) *Factory {
	switch componentType {
	case "receivers":
		return receivers[name]
	case "processors":
		return processors[name]
	case "exporters":
		return exporters[name]
	}
	panic("x")
}
