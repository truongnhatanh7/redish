package main

import (
	"fmt"
	"log"
	"net"
)

type Peer struct {
	conn    net.Conn
	msgChan chan []byte
}

func NewPeer(conn net.Conn, msg chan []byte) *Peer {
	return &Peer{
		conn:    conn,
		msgChan: msg,
	}
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			log.Println("peer read error", "err", err)
			return err
		}
		fmt.Println(string(buf[:n]))
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgChan <- msgBuf
	}
}
