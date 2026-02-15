package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Result struct {
	Port    int
	Open    bool
	Service string
}

func scanner(ports <-chan int, results chan<- Result, host string, wg *sync.WaitGroup) {
	defer wg.Done()

	dialer := net.Dialer{
		Timeout:   2 * time.Second,
		KeepAlive: 0,
	}

	for port := range ports {
		address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

		conn, err := dialer.Dial("tcp", address)
		if err != nil {
			continue
		}

		service := grabService(conn)

		results <- Result{
			Port:    port,
			Open:    true,
			Service: service,
		}

		conn.Close()

	}
}

func grabService(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}

	return string(buffer[:n])
}

func main() {
	host := "rupp.edu.kh"

	const maxWorker = 1000
	const maxPorts = 65535

	ports := make(chan int, 100)
	results := make(chan Result)
	var wg sync.WaitGroup

	fmt.Println("Scanning host:", host)
	// start workers
	for i := 1; i <= maxWorker; i++ {
		wg.Add(1)
		go scanner(ports, results, host, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for port := 1; port <= maxPorts; port++ {
			ports <- port
		}
		close(ports)
	}()

	for result := range results {
		fmt.Printf("Port %d is open.", result.Port)
		if result.Service != "" {
			fmt.Printf(" Service: %s", result.Service)
		} else {
			fmt.Println()
		}
	}
}
