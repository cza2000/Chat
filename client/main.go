package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "121.36.4.182:8000")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client begin")
	checkOrRegister(conn)

	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			return
		} // NOTE: ignoring errors
		log.Println("done")
		err2 := conn.(*net.TCPConn).CloseRead()
		if err2 != nil {
			return
		}
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	err = conn.(*net.TCPConn).CloseWrite()
	if err != nil {
		return
	}
	<-done // wait for background goroutine to finish

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}


