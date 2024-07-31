package channel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var firstSigusr1 = true

func RunChan() {
	// 忽略信号, 内核将中断丢弃, 中断信号对进程不产生影响
	signal.Ignore(os.Interrupt)

	// 有缓冲区的chan, chan的类型是os.Signal.
	// 因为 signal 包不会为了向 c 发送信息而阻塞（就是说如果发送时 c 阻塞了，signal 包会直接放弃）：
	// 调用者应该保证 c 有足够的缓存空间可以跟上期望的信号频率。
	// 对使用单一信号用于通知的 channel，缓存为 1/2 就足够了。
	c1 := make(chan os.Signal, 2)

	// 改变信号的默认行为, 会将SIGHUP信号转发到chan c1中
	signal.Notify(c1, syscall.SIGHUP)

	signal.Notify(c1, syscall.SIGUSR1)

	go func() {
		for {
			switch <-c1 {
			case syscall.SIGHUP:
				fmt.Println("SIGHUP, then reset SIGHUP")
				signal.Reset(syscall.SIGHUP)
			case syscall.SIGUSR1:
				if firstSigusr1 {
					fmt.Println("first usr1, notify interrupt which had ignore!")
					c2 := make(chan os.Signal, 1)
					signal.Notify(c2, os.Interrupt)
					go handlerInterrupt(c2)
				}
			}
		}
	}()

	select {}
}

func handlerInterrupt(c <-chan os.Signal) {
	for {
		switch <-c {
		case os.Interrupt:
			fmt.Println("signal interrupt")
		}
	}
}
