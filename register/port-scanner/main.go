package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const ip = ""
const maxPort = 65535

func scan(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return
	}
	defer conn.Close()

	fmt.Println(target)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < maxPort; i++ {
		wg.Add(1)
		go scan(i, &wg)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
}
