package main

import (
	"context"
	"fmt"
	"time"
)

var count int = 0

func main() {
	product := make(chan int, 10)
	ctx, cancel := context.WithCancel(context.TODO())

	for i := 0; i < 10; i++ {
		go Producter(product, ctx)
		go Consumer(product, ctx, cancel)
	}
	<-ctx.Done()
}

func Producter(ch chan<- int, ctx context.Context) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产了产品", count)
		ch <- count
		count++
	}
	// 生产结束了
	if count == 100 {
		close(ch)
	}
	return
}

func Consumer(ch <-chan int, ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		select {
		case v, ok := <-ch:
			// ch还没关闭，也就是生产活动还没结束
			if ok {
				fmt.Printf("消费者消费了产品%d\n ", v)
			} else { // ch 已经关闭了，以后都不会再生产了
				fmt.Println("所有生产的商品都已经消费完了")
				cancelFunc()
				// 想从 for select 中跳出，break是不行的，只能用goto跳出
				goto END
			}
		default:
			fmt.Printf("目前没有产品了，正在加紧生产中\n")
			time.Sleep(time.Second)
		}
	}
END:
	return
}
