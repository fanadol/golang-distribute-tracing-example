package tracing

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

// TraceMsg will be used as an io.Writer and io.Reader for the span's context and
// the payload. The span will have to be written first and read first.
// https://github.com/nats-io/not.go/blob/master/not.go
type TraceMsg struct {
	bytes.Buffer
}

// NewTraceMsg creates a trace msg from a NATS message's data payload.
func NewTraceMsg(m *stan.Msg) *TraceMsg {
	b := bytes.NewBuffer(m.Data)
	return &TraceMsg{*b}
}
