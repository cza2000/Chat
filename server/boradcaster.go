package main

import (
	"fmt"
	"strings"
)

type client chan<- string // an outgoing message channel

type clientInfo struct {
	ch client
	name string
	password string
	isOnline bool
}

var (
	entering = make(chan loginInfo)
	leaving  = make(chan loginInfo)
	messages = make(chan string) // all incoming client messages
	clients = make(map[string]*clientInfo) // all connected clients
)

func broadcaster() {

	for {
		select {
		case msg := <-messages://说话
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			fmt.Println(msg)
			for _, cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering://登录
			var users []string
			for cliName := range clients {
				if clients[cliName].isOnline {
					users = append(users, cliName)
				}
			}
			if len(users) > 0 {
				cli.ch <- fmt.Sprintf("Other users: %s", strings.Join(users, ","))
			} else {
				cli.ch <- "You are the only user."
			}
			clients[cli.name].ch = cli.ch
			clients[cli.name].isOnline = true

		case cli := <-leaving://注销
			clients[cli.name].isOnline = false
			clients[cli.name].ch = nil
			close(cli.ch)
		}
	}
}
