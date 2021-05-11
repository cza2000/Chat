package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client begin")
	checkOrRegister(conn)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		conn.(*net.TCPConn).CloseRead()
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done // wait for background goroutine to finish

}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func checkOrRegister(conn net.Conn) {

	fmt.Println("register(0) or login(1):")

	var status string
	var name string
	var password string
	fmt.Scanln(&status)
	buf := make([]byte, 4)
	switch status {
	case "0": //注册
		for {
			fmt.Println("Please input your name:")

			fmt.Scanln(&name)
			conn.Write([]byte("RegisterID:" + name))
			n, err := conn.Read(buf) //阻塞

			if err != nil {
				fmt.Println("conn read error:", err)
				return
			}
			//fmt.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
			if string(buf[:n]) == "1" {
				break
			}

			fmt.Printf("name %q is existed\r\ntry other name: \n", name)
		}

		for {
			fmt.Println("Please input your password:")

			fmt.Scanln(&password)
			conn.Write([]byte("RegisterPassword:" + password))
			break
		}

	case "1": //检查
		for {
			fmt.Println("Please input your name:")

			fmt.Scanln(&name)
			conn.Write([]byte("UserID:" + name))

			n, err := conn.Read(buf)

			if err != nil {
				fmt.Println("conn read error:", err)
				return
			}
			//fmt.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))

			if string(buf[:n]) == "1" {
				break
			}

			fmt.Printf("name %q is not existed\r\nplease input again: ", name)
		}

		for {
			fmt.Println("Please input your password:")

			fmt.Scanln(&password)
			conn.Write([]byte("UserPassword:" + password))
			n, err := conn.Read(buf)

			if err != nil {
				fmt.Println("conn read error:", err)
				return
			}
			//fmt.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))

			if string(buf[:n]) == "1" {
				break
			}
			fmt.Println("wrong password\r\nplease input again: ")
		}
	}

	return
}
