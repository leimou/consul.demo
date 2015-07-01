package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// Number of connections clients will create
var conns int

func init() {
	flag.IntVar(&conns, "c", 0, "Number of connections")
}

type ClientFEP struct {
	waitGroup sync.WaitGroup
}

func NewClient() *ClientFEP {
	return &ClientFEP{}
}

func (c *ClientFEP) Connect(id int) error {
	defer c.waitGroup.Done()

	conn, err := net.Dial("tcp4", "192.168.16.202:5000")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Connection established:", id)

	buf := make([]byte, 32*1024)
	for {
		_, err := conn.Write([]byte("Ping"))
		if err != nil {
			fmt.Println("conn write", id, ":", err.Error())
			return err
		}

		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn read", id, ":", err.Error())
			return err
		} else {
			fmt.Printf("conn %d: received %s\n", id, string(buf))
		}
		time.Sleep(time.Second)
	}
}

func main() {
	flag.Parse()

	c := NewClient()
	for i := 0; i < conns; i++ {
		c.waitGroup.Add(1)
		go c.Connect(i)
	}
	c.waitGroup.Wait()
}
