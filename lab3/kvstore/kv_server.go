// +build !solution

// Leave an empty line above this comment.
package main

import (
	"net"
	"net/rpc"
	"sync"
)
//
func NewKVStore(srvAddr string) {
	kv := &KVStore{new(sync.Mutex), make(map[string]string)}
	rpc.Register(kv)
	tcpAddr, err := net.ResolveTCPAddr("tcp",srvAddr)
	 checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	 checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

