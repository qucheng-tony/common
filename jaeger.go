package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegerCfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	return tracer, closer, err
}
