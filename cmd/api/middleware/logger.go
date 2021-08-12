package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		wi := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = wi
		c.Next()

		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.RequestURI,
			"status": c.Writer.Status(),
		}).Info(wi.body.String())
	}
}
