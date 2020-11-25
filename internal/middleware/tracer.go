package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"gotour/blog-service/global"
)

// Traceing 链路跟踪
func Traceing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				})
		}

		defer span.Finish()
		c.Request = c.Request.WithContext(newCtx)

		// 在 context 中保存traceId和spanId，便于日志记录
		setTraceId(c, span)

		c.Next()
	}
}

func setTraceId(c *gin.Context, span opentracing.Span) {

	var traceId string
	var spanId string
	var spanContext = span.Context()

	switch spanContext.(type) {
	case jaeger.SpanContext:
		jaegerContext := spanContext.(jaeger.SpanContext)
		traceId = jaegerContext.TraceID().String()
		spanId = jaegerContext.SpanID().String()
	}

	c.Set("X-Trace-ID", traceId)
	c.Set("X-Span-ID", spanId)
}
