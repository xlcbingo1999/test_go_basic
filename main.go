package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/xlcbingo1999/test_go_basic/cmd"
)

func runpprof(fn func() error) {
	srv := http.Server{
		Addr: ":6060",
	}

	go func() {
		fmt.Println("start in http://localhost:6060/debug/pprof")
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	fn()

	// 关闭http server
	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server shutdown gracefully")
}

func main() {
	// runpprof(cmd.Execute)
	cmd.Execute()
}
