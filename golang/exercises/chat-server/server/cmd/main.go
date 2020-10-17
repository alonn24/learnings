package main

import (
	"chat-server/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func createTcpListener(address string) net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Print(err)
		return nil
	}
	return listener
}

func onInterrupt(fn func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func main() {
	s := server.NewTcpServer()
	onInterrupt(func() { s.Close() })
	listener := createTcpListener(":3333")
	if listener != nil {
		log.Printf("Starting server on :3333")
		s.Listen(listener)
	}
}
