package main

import (
	"fmt"
	"time"
)

func main() {
	product := make(chan int, 10)
	wait := make(chan struct{})
	go Producter(product)
	go Consumer(product, wait)
	// 等待 wait 里有人发送了数据，不然一直阻塞主 goroutine 的退出
	// 从而达到不需要 time.Sleep()就能及时退出的效果
	<-wait
}

func Producter(ch chan<- int) {
	// 这里把产品全部发送之后，通道需要关闭吗？
	// 如果需要关闭，在哪里关闭呢？
	for i := 0; i < 100; i++ {
		fmt.Println("生产了产品", i)
		ch <- i
	}
	close(ch)
	return
}

func Consumer(ch <-chan int, wait chan<- struct{}) {
	// wait 是要等消费者给他发消息的，那么在这个函数里要在哪里发呢
	// 试着填空吧
	for {
		select {
		case v, ok := <-ch:
			// TODO: 如果ch还没关闭，也就是生产活动还没结束
			if ok {
				fmt.Printf("消费者消费了产品%d\n ", v)
			} else { // ch 已经关闭了，以后都不会再生产了
				fmt.Println("所有生产的商品都已经消费完了")
				wait <- struct{}{}
			}
		default:
			// default 的意义是什么呢？请尝试把你的理解打印输出。
			fmt.Printf("没有通道有可接收的值了，我要退出了\n")
			time.Sleep(time.Second)
		}
	}
}
