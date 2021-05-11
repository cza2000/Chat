package main

import (
	"fmt"
	"net"
	"strings"
)

func handleCheckOrRegister(conn net.Conn) (name, password string) {

	buf := make([]byte, 30)
	label:	for {
		n, _ := conn.Read(buf)
		request := strings.Split(string(buf[:n]), ":")

		switch request[0] {
		case "RegisterID":
			fmt.Println(request[0] + ", " + request[1])
			name = request[1]
			_, ok := clients[name]
			if ok {
				conn.Write([]byte("0"))
			} else {
				conn.Write([]byte("1")) //阻塞
				newRegisterInfo := clientInfo{
					ch:       nil,
					name:     name,
					password: "",
					isOnline: false,
				}
				clients[name] = &newRegisterInfo
			}

		case "RegisterPassword":
			fmt.Println(request[0] + ", " + request[1])
			password = request[1]
			clients[name].password = password
			break label

		case "UserID":
			fmt.Println(request[0] + ", " + request[1])
			name = request[1]
			_, ok := clients[name]
			if ok {
				conn.Write([]byte("1"))
			} else {
				conn.Write([]byte("0"))
			}

		case "UserPassword":
			fmt.Println(request[0] + ", " + request[1])
			password = request[1]
			if password == clients[name].password {
				conn.Write([]byte("1"))
			} else {
				conn.Write([]byte("0"))
			}
			break label
		}
	}
	return
}


