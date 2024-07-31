package contextX

import (
	"context"
	"log"
	"net/http"
)

type requestIDKeyType int

const requestIDKey requestIDKeyType = 0

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			// 从header中获取request-id
			reqID := req.Header.Get("X-Request-ID")
			// 创建一个context, 可以将自定义的数据传出去, 然后给其他routine接收
			// 这里是基于HTTP Request的context再产生的一个新的context
			// key是requestIDKey, value是reqID
			ctx := context.WithValue(
				req.Context(), requestIDKey, reqID,
			)

			// 创建新的请求
			req = req.WithContext(ctx)

			// 调用HTTP处理函数, 这里其实是外部注入的Handle, 用于打印内容
			next.ServeHTTP(rw, req)
		},
	)
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	reqID := GetRequestID(req.Context()) // 此时从req的上下文中就可以读到新建的ctx写入的requestIDKey的对应的值了
	log.Println("xlc get reqID: ", reqID)
}

func RunContextExampleHTTP() {
	// 注册了HTTP的一个路由方法
	handler := WithRequestID(http.HandlerFunc(Handle))

	// 测试方法: curl --noproxy "*" -v -H "Connection: close" -H "X-Request-ID: xlc123" localhost:8080
	http.ListenAndServe(":8080", handler)
}
