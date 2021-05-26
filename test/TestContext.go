package main

import (
	"context"
	"fmt"
	"time"
)

type Options struct {
	Interval time.Duration
}

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
			op.Interval = 2
		}
	}
}

func main() {
	// 创建可取消的 Context 对象，即可以主动通知子协程退出
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)                  // 超时退出
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(8*time.Second)) // 最迟退出时间
	vCtx := context.WithValue(ctx, "options", &Options{3})                                   // 往子协程中传递参数
	fmt.Println(ctx.Deadline())
	// 控制单个或多个协程
	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")

	time.Sleep(7 * time.Second)
	fmt.Println("before cancel")
	cancel()
	fmt.Println(vCtx.Value("options").(*Options))
	time.Sleep(3 * time.Second)
}
