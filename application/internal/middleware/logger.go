package middleware

import (
	"bytes"
	"fmt"
	"github.com/Jadepypy/distributed-social-media/application/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reflect"
	"time"
)

type LogMiddlewareBuilder struct {
	logger        logger.Logger
	allowReqBody  bool
	allowRespBody bool
}

type AccessLog struct {
	Path     string        `json:"path"`
	Method   string        `json:"method"`
	ReqBody  string        `json:"req_body"`
	Status   int           `json:"status"`
	RespBody string        `json:"resp_body"`
	Duration time.Duration `json:"duration"`
}

func NewLogMiddlewareBuilder(l logger.Logger) *LogMiddlewareBuilder {
	return &LogMiddlewareBuilder{
		logger: l,
	}
}

func (l *LogMiddlewareBuilder) AllowReqBody() *LogMiddlewareBuilder {
	l.allowReqBody = true
	return l
}

func (l *LogMiddlewareBuilder) AllowRespBody() *LogMiddlewareBuilder {
	l.allowRespBody = true
	return l
}

func (l *LogMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if len(path) > 1024 {
			path = path[:1024]
		}
		method := ctx.Request.Method
		al := AccessLog{
			Path:   path,
			Method: method,
		}
		if l.allowReqBody {
			body, _ := ctx.GetRawData()
			if len(body) > 2048 {
				al.ReqBody = string(body[:2048])
			} else {
				al.ReqBody = string(body)
			}
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		start := time.Now()

		if l.allowRespBody {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				al:             &al,
			}
		}

		defer func() {
			al.Duration = time.Since(start)
			l.handleError(ctx)
			l.logger.Debug("", logger.Field{Key: "req", Value: al})
		}()

		ctx.Next()
	}
}

func (l *LogMiddlewareBuilder) handleError(ctx *gin.Context) {
	for index, e := range ctx.Errors {
		err := e.Err
		if myErr, ok := err.(*PrivateError); ok {
			// Hide the private error reason from user
			// always return 200
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": myErr.Message,
			})
			// Replace the error with private error
			ctx.Errors[index] = &gin.Error{
				Err:  myErr,
				Type: gin.ErrorTypePrivate,
			}
			l.logger.Error("", logger.Field{Key: "Error", Value: myErr})
		} else {
			// return 500 if not specify any error code
			code := http.StatusInternalServerError
			elem := reflect.ValueOf(err).Elem()
			fmt.Printf("elem: %v\n", elem)
			field := elem.FieldByName("Code")
			if field.IsValid() {
				code = int(field.Int())
			}

			ctx.JSON(code,
				PrivateError{
					code,
					GeneralResponse{
						"Server Internal Error",
						err.Error(),
					},
				},
			)
			l.logger.Error("", logger.Field{Key: "Error", Value: err})
		}
		return
	}
}

type GeneralResponse struct {
	Message string `json:"message"` // returned to user
	Reason  string `json:"reason"`  // error reason if any (internal display only)
}

type PrivateError struct {
	Code int `json:"code"`
	GeneralResponse
}

func (e *PrivateError) Error() string {
	return e.Reason
}

type responseWriter struct {
	gin.ResponseWriter
	al *AccessLog
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.al.RespBody = string(data)
	return w.ResponseWriter.Write(data)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.al.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
