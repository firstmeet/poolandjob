package job

import (
	"awesomeProject/internal"
	"fmt"
)

var Job chan *internal.Package
var Stop = make(chan struct{})
var Continue = make(chan struct{})
var Number int64 = 0

func NewJob() {
	Job = make(chan *internal.Package, 100)
}
func DoJob() {
	go func() {
		for {
			select {
			case pkg := <-Job:
				handleJob[pkg.Type](pkg.Body)
			case <-Stop:
				<-Continue
				fmt.Println("continue job")
			}
		}
	}()

}
