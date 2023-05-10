package log

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"os"
	"reflect"
)

type Wrapper struct {
	log    Interface
	fields map[string]interface{}
}

func NewWrapper(l Interface) *Wrapper {
	return &Wrapper{
		log:    l,
		fields: map[string]interface{}{},
	}
}

func (w *Wrapper) Trace(args ...interface{}) {
	if !w.log.Options().level.Enabled(TraceLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(DebugLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(TraceLevel, args...)
}

func (w *Wrapper) Debug(args ...interface{}) {
	if !w.log.Options().level.Enabled(DebugLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(DebugLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(DebugLevel, args...)
}

func (w *Wrapper) Info(args ...interface{}) {
	if !w.log.Options().level.Enabled(InfoLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(InfoLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(InfoLevel, args...)
}

func (w *Wrapper) Warn(args ...interface{}) {
	if !w.log.Options().level.Enabled(WarnLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(WarnLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(WarnLevel, args...)
}

func (w *Wrapper) Error(args ...interface{}) {
	if !w.log.Options().level.Enabled(ErrorLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(ErrorLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(ErrorLevel, args...)
}

func (w *Wrapper) Fatal(args ...interface{}) {
	if !w.log.Options().level.Enabled(FatalLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
		ns[LogLineNumKey] = fileWithLineNum(w.log.Options())
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(FatalLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(FatalLevel, args...)
	os.Exit(1)
}

func (w *Wrapper) WithError(err error) *Wrapper {
	ns := copyFields(w.fields)
	ns[LogErrorKey] = err
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func (w *Wrapper) WithFields(fields map[string]interface{}) *Wrapper {
	ns := copyFields(fields)
	for k, v := range w.fields {
		ns[k] = v
	}
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func (w *Wrapper) WithContext(ctx context.Context) *Wrapper {
	requestId, traceId, spanId := GetId(ctx)
	if requestId == "" {
		return w
	}
	ns := copyFields(w.fields)
	if traceId != "" {
		ns[MiddlewareTraceIdCtxKey] = traceId
		ns[MiddlewareSpanIdCtxKey] = spanId
	} else {
		ns[MiddlewareRequestIdCtxKey] = requestId
	}
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func RequestId(ctx context.Context) (id string) {
	ctx = RealCtx(ctx)
	// get value from context
	requestIdValue := ctx.Value(MiddlewareRequestIdCtxKey)
	if item, ok := requestIdValue.(string); ok && item != "" {
		id = item
	}
	return
}

func GetId(ctx context.Context) (string, string, string) {
	ctx = RealCtx(ctx)
	requestId := RequestId(ctx)
	traceId, spanId := TraceId(ctx)
	if traceId != "" {
		requestId = traceId
	}
	return requestId, traceId, spanId
}

func TraceId(ctx context.Context) (traceId, spanId string) {
	ctx = RealCtx(ctx)
	span := trace.SpanFromContext(ctx).SpanContext()
	if span.IsValid() {
		traceId = span.TraceID().String()
		spanId = span.SpanID().String()
	}
	return
}

func RealCtx(ctx context.Context) context.Context {
	if interfaceIsNil(ctx) {
		return context.Background()
	}
	if c, ok := ctx.(*gin.Context); ok {
		// gin context contains cancel ctx, remove it
		ctx = c.Request.Context()
		requestId, traceId, spanId := GetId(ctx)
		if traceId != "" {
			ctx = context.WithValue(ctx, MiddlewareTraceIdCtxKey, traceId)
			ctx = context.WithValue(ctx, MiddlewareSpanIdCtxKey, spanId)
		} else {
			ctx = context.WithValue(ctx, MiddlewareRequestIdCtxKey, requestId)
		}
	}
	return ctx
}

func interfaceIsNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return i == nil
}
