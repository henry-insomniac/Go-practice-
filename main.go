// 第一个 Go
package main

import (
	"fmt"
	"time"

	"github.com/henry-insomniac/Go-practice-/localpackage"
)

func main() {
	fmt.Println("hello Go")
	// 每隔1秒执行一次
	t := time.Tick(1 * time.Second)
	for {
		<-t
		currentTime := localpackage.GetCurrentTime()
		fmt.Println("当前时间：", currentTime)
	}
}
