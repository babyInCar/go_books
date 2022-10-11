// This is a program about goroutine and fibnacci!
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibbnacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// tips： 在这个程序中，使用go关键字来另外起了一个goroutine来在输出结果之前
// 给出一些提示符，这种对于运行时间比较久的程序是一种友好的提示
// 同时也发挥了go语言并发执行的优势
