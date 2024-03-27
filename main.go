package main

import "github.com/xlcbingo1999/test_go_basic/cmd"

func main() {
	cmd.Execute()

	ch := make(chan int, 2)
	ch <- 2
}
