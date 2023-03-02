package job

import (
	"awesomeProject/internal"
	"fmt"
	"sync/atomic"
)

var handleJob map[int]func([]byte) = make(map[int]func([]byte))

func RegisterJobHandleFunc(jobType int, fn func([]byte)) {
	handleJob[jobType] = fn
}
func init() {
	RegisterJobHandleFunc(internal.StartJob, startJob)
	RegisterJobHandleFunc(internal.StopJob, stopJob)
}
func startJob(body []byte) {
	atomic.AddInt64(&Number, 1)
	fmt.Println(string(body))
	fmt.Printf("已完成%d个任务\n", atomic.LoadInt64(&Number))
}
func stopJob(body []byte) {
	fmt.Println("stop job")
	go func() {
		Stop <- struct{}{}
	}()
}
func ContinueJob() {
	go func() {
		Continue <- struct{}{}
		fmt.Println("continue job")
	}()

}
