package Struct

import "fmt"

// 最常用的设计模式: 本质上就是构建一个真实服务的代替品给客户端使用, 代理和真实服务需要使用同样的接口
// 这里的例子是用Nginx做一个反向代理, 有很多种真实服务的实现, 最后需要确定服务要转发到哪些上面去
// Nginx需要设置最大的请求量, 如果一个url已经达到某个请求量, 就不再支持了
// rateLimiter 用于做一些负载均衡策略, key是每个url, value是每个url对应的请求量

type server interface {
	handleRequest(string, string) (int, string)
}

type Nginx struct {
	application     *Application
	maxAllowRequest int
	rateLimiter     map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		application:     &Application{},
		maxAllowRequest: 2,
		rateLimiter:     make(map[string]int),
	}
}

// Nginx也需要实现server的接口
func (n *Nginx) handleRequest(url string, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	time, ok := n.rateLimiter[url]
	if ok {
		if time >= n.maxAllowRequest {
			return false
		} else {
			n.rateLimiter[url] += 1
		}
	} else {
		n.rateLimiter[url] = 1
	}
	return true
}

type Application struct {
}

func (a *Application) handleRequest(url string, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "Created"
	}

	return 404, "Failed"
}

func RunProxy() {
	nginxServer := NewNginx()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
