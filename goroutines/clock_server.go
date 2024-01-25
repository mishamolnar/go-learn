package goroutines

import (
	"io"
	"log"
	"net"
	"time"
)

func ClockServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleClockConn(conn)
	}
}

func handleClockConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05:123\n"))
		if err != nil {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
}
