package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 假设商品总量为2
var count int = 2

// 假设中奖率为 2%
var rate int = 2

// 这里我们的out需要带缓冲区吗
var out chan int = make(chan int)

var wg sync.WaitGroup

// 模拟一次抽奖请求
func Get(id int) <-chan int {
	lucky := rand.Int() % 100
	if lucky < rate {
		out <- id
	}
	return out
}

// 模拟发放奖品
func Prize(ch <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case value, _ := <-ch:
			if count > 0 {
				count--
				fmt.Println("id为:", value, "的用户获奖了")
			} else {
				// 为了方便查结果这里就不打印了
				fmt.Println("奖品已发放完毕")
				wg.Done()
				goto END
			}
		default:
			// 为了方便查结果这里就不打印了
			// fmt.Println("未获奖")
			time.Sleep(time.Second)
		}
	}
END:
	return
}

func main() {
	wg.Add(1)
	go Prize(out, &wg)
	for i := 0; i < 10000; i++ {
		go Get(i)
	}
	wg.Wait()
	defer close(out)

}
