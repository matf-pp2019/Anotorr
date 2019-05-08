package main

import (
	"./structs/IO"
	"./structs/Node"
	"fmt"
	"net"

	//"net"
)

var trackerReader = IO.Reader{nil}
var trackerWriter = IO.Writer{nil}

func main() {
	/*
	//Inicijalizaciju Merkle stabla isto prebaci u InitializeNode()
	m := MerkleTree.Merkle{make([][] string, 0)}

	m.CreateTree("misc/zorka.mp3", 5, 1000000)

	m.CreateProof(1)
	*/
	self := Node.InitializeNode()
	//fmt.Printf("[PeerMain] Hello, my name is: %+v\n", self)

	//Javljam se trekeru. Hardkodovan localhost
	tAddr, err := net.ResolveTCPAddr("tcp", "192.168.0.13:9095")
	Node.CheckError(err)
	conn, err := net.DialTCP("tcp",nil, tAddr)
	Node.CheckError(err)

	self.ReqConn = conn

	// Citac i pisac otvoreni ka trekeru za postavjanje requestova
	trackerReader = IO.Reader{self.ReqConn}
	trackerWriter = IO.Writer{self.ReqConn}

	// Poruka predstavljanja trekera, choose option itd...
	msg := trackerReader.Read()
	fmt.Println(msg)

	var ans string
	_, err = fmt.Scanf("%s", &ans)
	Node.CheckError(err)

	if (ans == "D") {
		self.RequestDownload(trackerWriter, trackerReader)
	}

	self.WaitGroup.Add(2)

		go self.ListenTracker()

		go self.ListenPeer()

	self.WaitGroup.Wait()
}