package contextX

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	KEY = "trace_id"
)

type myString string

func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	reqID := NewRequestID()
	ctx := context.WithValue(context.Background(), myString(KEY), reqID)
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "xlc")
}

func RunContextExample2() {
	ProcessEnter(NewContextWithTraceID())
}
