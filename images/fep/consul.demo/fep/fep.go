package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync/atomic"
)

type ServiceFEP struct {
	// Number of accepted connections
	conns int32
}

func NewService() *ServiceFEP {
	return &ServiceFEP{
		conns: 0,
	}
}

func (fep *ServiceFEP) Loop(conn net.Conn) error {
	defer func() {
		atomic.AddInt32(&fep.conns, -1)
	}()

	buf := make([]byte, 32*1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
}

func (fep *ServiceFEP) Serve() error {
	l, err := net.Listen("tcp4", ":32768")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		atomic.AddInt32(&fep.conns, 1)
		log.Println("Accepted from:", conn.RemoteAddr().String())
		go fep.Loop(conn)
	}
}

func (fep *ServiceFEP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%d", atomic.LoadInt32(&fep.conns))
}

func main() {
	fep := NewService()
	go fep.Serve()

	http.Handle("/conns", fep)
	http.ListenAndServe(":5000", nil)
}
