package noMutexDemo

import (
	"log"
	"net/http"
	"time"

	_ "net/http/pprof" // 自动增加的性能分析插件

	"github.com/xlcbingo1999/test_go_basic/data"
)

func RunNoMutexDemo() {
	go func() {
		for {
			log.Println(data.Add("github.com/xlcbingo1999"))
			time.Sleep(time.Second * 2)
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}
