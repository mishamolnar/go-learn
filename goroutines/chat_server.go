package goroutines

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Chat() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnectionChat(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for clt := range clients {
				clt <- msg
			}
		case clt := <-entering:
			clients[clt] = true
		case clt := <-leaving:
			delete(clients, clt)
			close(clt)
		}
	}
}

func handleConnectionChat(conn net.Conn) {
	ch := make(chan string) //outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "Has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " left the chat"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for line := range ch {
		fmt.Fprintln(conn, line)
	}
}
