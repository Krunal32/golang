// +build !solution
// Leave an empty line above this comment.
//
// Zap Collection Server
package main

import (
	"fmt"
	"github.com/uis-dat320-fall2014/labs/zaplab/chzap"
	"github.com/uis-dat320-fall2014/labs/zaplab/zlog"
	"net"
	"os"
	"time"
)
var latestzap chan chzap.ChZap
var isRunning bool = true
var multicastAddr, rpcServerAddr string = "224.0.1.130:10000", ":12110"

func RunLab() {
	latestzap = make(chan chzap.ChZap)
	switch *labnum {
	case "a", "c1", "c2":
		ztore = zlog.NewSimpleZapLogger()
	case "e", "f", "g":
		ztore = zlog.NewViewersZapLogger()
	}
	switch *labnum {
	case "a":
		go A_DumpToConsole(10)
	case "c1":
		go C1_C2_ShowViewers("NRK1", 1)
	case "c2":
		go C1_C2_ShowViewers("TV2 Norge", 1)
	case "e":
		go E_PrintTop10(3)
	case "f":
		go RunRPCServer(&ztore, rpcServerAddr) // start rpc server
	case "g": // test AverageDurations
		go G_PrintAverageDurations(100)
	}
}
func StartServer() {
	addrMulticastUDP, err := net.ResolveUDPAddr("udp", multicastAddr) // get address
	checkError(err, "Error resolving")
	conn, err := net.ListenMulticastUDP("udp", nil, addrMulticastUDP) //connect to zap server
	checkError(err, "Error multicast")
	go readZaps(conn)

}

func A_DumpToConsole(del int) { // prints received zap events
	for _ = range time.Tick((time.Duration(del)) * time.Millisecond) { //delay for readability
		ch := <-latestzap //reads from channel
		fmt.Printf("%s\n", ch)
	}
}
func C1_C2_ShowViewers(channel string, del int) { // prints viewers for a specified channel
	for _ = range time.Tick((time.Duration(del)) * time.Second) {
		fmt.Printf("Viewers on %s : %d \n", channel, ztore.Viewers(channel))
	}
}
func E_PrintTop10(del int) { //top 10 channels

	for _ = range time.Tick((time.Duration(del)) * time.Second) {
		fmt.Println(ztore.Top10())
	}
}

func G_PrintAverageDurations(del int) {

	for _ = range time.Tick((time.Duration(del)) * time.Millisecond) { //delay for readability
		dur := ztore.AverageDuration()
		fmt.Printf("Average duration: %0.2fs\n", dur.Seconds())
	}
}

func readZaps(conn net.Conn) {
	var buf [1024]byte
	for isRunning {
		if n, err := conn.Read(buf[0:]); err == nil {
			str := string(buf[:n])
			chz, _ := chzap.NewSTBEvent(str) // create ChZap from string
			ztore.LogZap(*chz)               // add zap to storage
			if ztore != nil {                //fmt.Printf("Read strzap: %s", str)
				ztore.LogZap(*chz)          // add zap to storage
				if *labnum == "a" {
					latestzap <- *chz  // send zaps over channel for printout 
				}
			}

		}
	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Message: %s | Fatal error:  %s", msg, err.Error())
		os.Exit(1)
	}
}
