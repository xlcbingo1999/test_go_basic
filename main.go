package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/xlcbingo1999/test_go_basic/cmd"
)

func pprofWithCancel() {
	go func() {
		fmt.Println("start in http://localhost:6060/debug/pprof")
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			panic(err)
		}
	}()
}

func main() {
	pprofWithCancel()
	cmd.Execute()
}
