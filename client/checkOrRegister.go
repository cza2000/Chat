package main

import (
	"fmt"
	"net"
)

func checkOrRegister(conn net.Conn) {

	fmt.Println("register(0) or login(1):")

	var status string
	var name string
	var password string
	_, err := fmt.Scanln(&status)
	if err != nil {
		return
	}
	buf := make([]byte, 4)
	switch status {
	case "0": //注册
		for {
			fmt.Println("Please input your name:")
			_, err := fmt.Scanln(&name)
			if err != nil {
				return
			}
			_, err2 := conn.Write([]byte("RegisterID:" + name))
			if err2 != nil {
				return
			}
			n, err3 := conn.Read(buf)
			if err3 != nil {
				fmt.Println("conn read error:", err3)
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

			_, err := fmt.Scanln(&password)
			if err != nil {
				return
			}
			_, err2 := conn.Write([]byte("RegisterPassword:" + password))
			if err2 != nil {
				return
			}
			break
		}

	case "1": //检查
		for {
			fmt.Println("Please input your name:")

			_, err := fmt.Scanln(&name)
			if err != nil {
				return
			}
			_, err2 := conn.Write([]byte("UserID:" + name))
			if err2 != nil {
				return
			}

			n, err3 := conn.Read(buf)

			if err3 != nil {
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

			_, err := fmt.Scanln(&password)
			if err != nil {
				return
			}
			_, err2 := conn.Write([]byte("UserPassword:" + password))
			if err2 != nil {
				return
			}
			n, err3 := conn.Read(buf)

			if err3 != nil {
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

