package futuresX

import (
	"fmt"
	"net/http"

	"github.com/Allan-Jacobs/go-futures"
)

// HTTPGetAsync wraps http.Get into a future based api
func HTTPGetAsync(url string) futures.Future[*http.Response] {
	return futures.GoroutineFuture(func() (*http.Response, error) {
		return http.Get(url)
	})
}

func RunFuture3rdParty() {
	// run futures simultaneously and await aggregated results
	responses, err := futures.All(
		HTTPGetAsync("https://go.dev"),
		HTTPGetAsync("https://pkg.dev"),
	).Await()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, res := range responses {
		fmt.Println(res.Request.URL, res.Status)
	}
}
