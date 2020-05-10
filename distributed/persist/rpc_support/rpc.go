package rpc_support

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, server interface{}) error {
	rpc.Register(server)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept error :%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
