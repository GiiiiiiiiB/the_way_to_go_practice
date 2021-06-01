package main

import "time"
import "fmt"
import "context"

//type Context interface {
//	Deadline() (deadline time.Time, ok bool)
//	Done() <-chan struct{}
//	Err() error
//	Value(key interface{}) interface{}
//}

const DB_ADDRESS  = "db_address"
const CALCULATE_VALUE  = "calculate_value"

func readDB(ctx context.Context, cost time.Duration)  {
	fmt.Println("db address is", ctx.Value(DB_ADDRESS))
	select {
	case <- time.After(cost): //  模拟数据库读取
		fmt.Println("read data from db")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 任务取消的原因
		fmt.Println("cuowuxinxi")
		// 一些清理工作
	}
}
func calculate(ctx context.Context, cost time.Duration)  {
	fmt.Println("calculate value is", ctx.Value(CALCULATE_VALUE))
	select {
	case <- time.After(cost): //  模拟数据计算
		fmt.Println("calculate finish")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 任务取消的原因
		fmt.Println("cuowuxinxi")
		// 一些清理工作
	}
}


func worker(done chan bool) {
	time.Sleep(3*time.Second)
	// 通知任务已完成
	done <- true
}



func main()  {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	fmt.Println(stop2)
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		i := 1
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			i++
			if i > 2 {
				ticker.Stop()
			}
		}
	}()

	time.Sleep(3*time.Second)

	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	<-done

	ctx := context.Background()
	fmt.Println("start test ..")
	fmt.Println(ctx)
	fmt.Println("stop test ..")
	ctx = context.WithValue(ctx,DB_ADDRESS,"nihaowa")
	ctx = context.WithValue(ctx,CALCULATE_VALUE,1234)

	ctx,cancel := context.WithTimeout(ctx,12* time.Second)
	defer cancel()

	go readDB(ctx,3*time.Second)
	go calculate(ctx,4*time.Second)

	time.Sleep(5*time.Second)
}

