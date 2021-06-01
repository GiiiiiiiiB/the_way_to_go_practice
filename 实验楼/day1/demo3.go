package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 准备好几个通道。
	Channels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	fmt.Printf("%d\n", index)
	// 向随机选择的通道里发送值
	Channels[index] <- index

	go func(Channels [3]chan int) {
		for {
			index := rand.Intn(3)
			fmt.Printf("%d\n", index)
			Channels[index] <- index
			time.Sleep(time.Second * 2)
		}

	}(Channels)
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	go func(Channels [3]chan int) {
	LOOP:
		for {
			select {
			case <-Channels[0]:
				fmt.Println("第一个被选中.")
			case <-Channels[1]:
				fmt.Println("第二个被选中")
			case elem := <-Channels[2]:
				fmt.Printf("第三个被选中，且值为：%d.\n", elem)
			default:
				fmt.Println("没有通道有可接收的值了，我要退出了")
				// 	// 想从 for select 语句中跳出，必须使用 goto 语句
				time.Sleep(time.Second)
				goto LOOP
			}
		}
	}(Channels)
	time.Sleep(time.Second * 20)
	// END:
	// 	return
}
