package deferx

import (
	"fmt"
	"log"
)

func RunDefer() {
	defer fmt.Println("in main")
	defer fmt.Println("second in main")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover: %v", err)
		}
		panic("panic second")
	}()

	panic("panic once")
}
