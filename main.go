package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"github.com/tomhjx/go-rsyslog-elastic-apm/input"
)

func main() {
	var tcpIface string
	flag.StringVar(&tcpIface, "tcp host:port", ":8080", "a tcp host:port")
	flag.Parse()
	host, err := os.Hostname()
	if err != nil {
		log.Fatal("unable to determine hostname")
	}
	fmt.Println("host is ", host)

	ch := make(chan string, 256)
	tcpServer = input.NewTcpServer(tcpIface)
	err = tcpServer.Up(ch)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("listening on %s for TCP connections", tcpIface)

	for {
		msg := <-ch
		log.Println("send data :", msg)
	}

}