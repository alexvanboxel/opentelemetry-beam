package googlecloudpubsubexporter

import (
	"bytes"
	"compress/gzip"
)

type compression int

const (
	uncompressed compression = iota
	gZip                     = iota
)

func compress(payload []byte, compression compression) ([]byte, error) {
	if compression == gZip {
		buf := new(bytes.Buffer)
		writer := gzip.NewWriter(buf)
		_, err := writer.Write(payload)
		if err != nil {
			return nil, err
		}
		err = writer.Close()
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return payload, nil
}

//const name =
//
//
//if (ceType != null) {
//switch (ceType) {
//case "org.opentelemetry.otlp.traces.v1":
//parseTrace(c, message);
//break;
//case "org.opentelemetry.otlp.metrics.v1":
//parseMetrics(c, message);
//break;
//case "org.opentelemetry.otlp.logs.v1":
//parseLogs(c, message);
//break;
//default:
//}
//}
