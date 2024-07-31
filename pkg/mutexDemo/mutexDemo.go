package mutexDemo

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var ServePort string

func init() {
	log.Println("start init mutex_demo")
	runtime.SetMutexProfileFraction(1)
}

func RunMutexDemo() {
	var m sync.Mutex
	var datas = make(map[int]struct{})

	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()

			datas[i] = struct{}{}
			time.Sleep(time.Second * 1)
		}(i)
	}

	log.Println("start serve in ", fmt.Sprintf("0.0.0.0:%s", ServePort))
	_ = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", ServePort), nil)

}
