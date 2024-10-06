package xconnector

import "net/http"

type Conn struct {
	http.Client
}

func NewConn() Conn {
	return Conn{}
}

func (c *Conn) Connect() {
}

func (c *Conn) Disconnect() {
}
