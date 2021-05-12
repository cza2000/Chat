package main

import (
	"bufio"
	"net"
)

type loginInfo struct {
	name string
	ch   client
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who, _ := handleCheckOrRegister(conn)

	ch <- "Login success:" + who
	cli := loginInfo {
		name: who,
		ch: ch,
	}
	entering <- cli

	done := make(chan struct{})
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			messages <- who + ": " + input.Text()
		}

		done <- struct{}{}
	}()

	<-done
	messages <- who + " has left"
	leaving <- cli

	conn.Close()
}
