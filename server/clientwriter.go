package main

import (
	"fmt"
	"net"
)

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		//fmt.Println(msg)
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors

	}
}