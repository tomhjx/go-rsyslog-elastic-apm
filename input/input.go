package input

import (
	"log"
	"net"
	"bufio"
	"time"
)

type server struct {
	iface string
}

type TcpServer struct {
	server
}

func NewTcpServer(iface string) *TcpServer {
	s := &TcpServer{}
	s.iface = iface

	return s
}

func (s *TcpServer) handleConn(conn net.Conn, ch chan<- string) {
	defer func() {
		log.Println("close a connection")
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		// conn.SetReadDeadline(time.Now().Add(1000))
		b, err := reader.ReadByte()
		if err != nil {
			log.Println("Error from connection:", err)
			return
		}
		time.Sleep(3)
		ch <- "ping"
		log.Println(string(b))
	}
}

func (s *TcpServer) Up(ch chan<- string) error {
	ln, err := net.Listen("tcp", s.iface)
	if err != nil {
		log.Fatal("failed to listen connection")
		return err
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("failed to accept connection", err)
				continue
			}
			log.Println("accepted new connection from", conn.RemoteAddr().String())
			go s.handleConn(conn, ch)
		}
	}()
	return nil
}