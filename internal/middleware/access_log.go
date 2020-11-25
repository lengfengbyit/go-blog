package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gotour/blog-service/global"
	"gotour/blog-service/pkg/helper"
	"gotour/blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// AccessLog 记录访问日志
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = bodyWriter

		beginTime := time.Now()
		c.Next()
		endTime := time.Now()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %s, end_time: %s, cost_time(ms): %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			helper.TimeFmt(beginTime, "datetime"),
			helper.TimeFmt(endTime, "datetime"),
			endTime.Sub(beginTime).Milliseconds(),
		)
	}
}
