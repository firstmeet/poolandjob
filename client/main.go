package main

import (
	"awesomeProject/internal"
	"awesomeProject/internal/network"
	"awesomeProject/internal/pool"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	var addr string
	var max uint64
	flag.StringVar(&addr, "l", "0.0.0.0:3000", "server address")
	flag.Uint64Var(&max, "m", 10, "max goroutine")
	flag.Parse()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	p := pool.New(max)
	defer p.Close()
	for i := 0; i < 100; i++ {
		p.Submit(func() {
			pkg := &internal.Package{Type: internal.StartJob, Body: []byte("hello")}
			err := network.SendPackage(conn, pkg)
			if err != nil {
				fmt.Println(err)
			}
		})
		if i == 3 {
			p.Submit(func() {
				pkg := &internal.Package{Type: internal.StopJob, Body: []byte("stop")}
				err := network.SendPackage(conn, pkg)
				if err != nil {
					fmt.Println(err)
				}
			})
		}

	}
	time.Sleep(time.Second * 3)
	pkg := &internal.Package{Type: internal.ContinueJob, Body: []byte("continue")}
	err = network.SendPackage(conn, pkg)
	if err != nil {
		fmt.Println(err)
	}
	select {}
}
