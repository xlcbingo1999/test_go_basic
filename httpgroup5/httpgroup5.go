package httpgroup5

import (
	"fmt"
	"sync"
)

func RunHttpGroup5() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println("Finished: ", i)
			wg.Done()
		}(i)
		if i > 0 && i%5 == 0 {
			fmt.Println("Stop wait")
			wg.Wait()
		}
	}
}
