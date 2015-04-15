package main

import (
	"fmt"
	"net"
"bufio" 
"os"
)

func clientLoop(service string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
    defer conn.Close()




fmt.Print("Enter: ")
    var buf [bsize]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Bytes()
		if string(s) == "exit" {
			// break out if user typed 'exit'
			break
		}
		_, err1 := conn.Write(s)

        checkError(err1)
	    n, err2 := conn.Read(buf[0:])
         checkError(err2)
	
		fmt.Println("Server: " + string(buf[0:n]))
		fmt.Print("Enter: ")
	}



}
