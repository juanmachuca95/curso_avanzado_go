package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("website", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < 6000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, i))
			if err != nil {
				return
			}

			conn.Close()
			fmt.Println("port is open: ", i)
		}(i)
	}

	wg.Wait()
}

// Ejecutar progama con flag : go run /net/port.go --website=scanme.webscantest.com
