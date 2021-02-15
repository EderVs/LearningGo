package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	c    chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.c <- msg
			}
		case cli := <-entering:
			sendParticipants(cli.c, clients)
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		}
	}
}

func sendParticipants(c chan<- string, clients map[client]bool) {
	allClients := ""
	for currentCli := range clients {
		allClients += fmt.Sprintf("%q, ", currentCli.name)
	}
	if len(clients) == 0 {
		allClients = "You are the only one in the chat."
	}
	if len(clients) == 1 {
		allClients += "is in the chat."
	} else if len(clients) >= 2 {
		allClients += "are in the chat."
	}
	c <- allClients
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	cli := createClient(ch, input)

	messages <- cli.name + " has arrived"
	entering <- cli

	finishCh := make(chan bool)
	continueCh := make(chan bool)
	connClosedCh := make(chan bool)
	go clientTimeout(finishCh, continueCh)
	go readMessages(input, cli.name, messages, continueCh, connClosedCh)

	select {
	case <-finishCh:
		fmt.Printf("timeout: %s\n", cli.name)
		conn.Close()
	case <-connClosedCh:
		fmt.Printf("connection closed: %s\n", cli.name)
	}
	leaving <- cli
	messages <- cli.name + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ingnoring network errors
	}
}

func createClient(ch chan string, input *bufio.Scanner) client {
	ch <- "Please enter your name: "
	input.Scan()
	who := input.Text()
	ch <- "Welcome, " + who
	return client{c: ch, name: who}
}

func clientTimeout(finishCh chan bool, continueCh chan bool) {
	for {
		select {
		case cont := <-continueCh:
			if !cont {
				return
			}
		case <-time.After(20 * time.Second):
			finishCh <- true
		}
	}
}

func readMessages(input *bufio.Scanner, who string, messages chan string, continueCh chan bool, connClosedCh chan bool) {
	for input.Scan() {
		continueCh <- true
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	// The connection is closed
	connClosedCh <- true
}
