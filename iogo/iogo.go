package iogo

import "runtime"

func RunIOgo() {
	runtime.GOMAXPROCS(1) // 这里就是GMP模型中, P的数量设置为1的场景

}
