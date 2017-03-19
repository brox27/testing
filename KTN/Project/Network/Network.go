// This is a server for KTN written in GOlang!

package Network

import (
//	. "fmt"
	"../ConfigFile"
	"encoding/json"
//	"time"
	"net"
)
const SV_LISTEN_ADDRESS = "127.0.0.1:12345"

func ServerTransmitter(sendchan chan ConfigFile.ResponseStruct){
	println("Transmitter Started...SERVER")
}

func ClientTransmitter(sendchan chan ConfigFile.Request){
	println("Transmitter Started...SERVER")
	for{
		
	}
}

func ClientListener(conn *net.TCPConn, RecieveChan chan ConfigFile.Request){
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)		//difference read og readfrom??
		if (err != nil){println("ERROR i ClientListener")}
		var NewReq ConfigFile.Request
		json.Unmarshal(buf[:n], NewReq)
		RecieveChan <- NewReq
	}
}

func FromServerListener(RecieveChan chan ConfigFile.ResponseStruct){
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)		//difference read og readfrom??
		if (err != nil){println("ERROR i ClientListener")}
		var NewReq ConfigFile.ResponseStruct
		json.Unmarshal(buf[:n], NewReq)
		RecieveChan <- NewReq
	}
}

func ConnectionListener(NewConnectionChan chan *net.TCPConn){
	println("ConnectionListener Started...")
	local, err := net.ResolveTCPAddr("tcp", SV_LISTEN_ADDRESS)

	ln, err := net.ListenTCP("tcp", local)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			// handle error
		}
		NewConnectionChan <- conn
	}
}
