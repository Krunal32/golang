// +build !solution

// Leave an empty line above this comment.
package main
import ("net"
     "fmt"
    )

func serverLoop(service string) {
	udpAddr, err := net.ResolveUDPAddr("udp",service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {

		handleClient(conn)

	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	for {
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.WriteToUDP(buf[0:n], addr)
		if err2 != nil {
			return
		}
	}
}

