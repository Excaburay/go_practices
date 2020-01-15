package learntest

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//calculate Fibonacci sequence
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Test1(n int) {
	go spinner(100 * time.Millisecond)
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func Test2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("create listener.Accept")
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//注意\n \r \t 的区别 mac \r  unix \n windows \n\r
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func Netcat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

//为了让程序等待后台的goroutine在完成后再退出
func Netcat3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //等待后台 goroutine 完成
}

func Pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)
	//自然数
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()
	//平方 当接收操作在一个关闭的且读完的通道上，ok为false
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()
	//主goroutine中打印
	for {
		fmt.Println(<-squares)
	}
}

//使用range方便判断
func Pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)
	//自然数
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	//平方
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	//主goroutine中打印
	for {
		fmt.Println(<-squares)
	}
}

//将上述过程拆分为多个函数
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Pipeline3() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
