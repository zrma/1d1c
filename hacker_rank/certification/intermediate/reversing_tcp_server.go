package intermediate

import (
	"context"
	"net"
	"strings"
)

const (
	maxBufferSize = 1024
	network       = "tcp"
)

func TCPServer(ctx context.Context, ready chan string) {
	const endpoint = ""
	addr, err := net.ResolveTCPAddr(network, endpoint)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP(network, addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer func() {
			_ = listener.Close()
		}()

		for ctx.Err() == nil {
			conn, err := listener.AcceptTCP()
			if err != nil {
				panic(err)
			}
			go handleConn(conn)
		}
	}()

	ss := strings.Split(listener.Addr().String(), ":")
	if len(ss) < 2 {
		ready <- listener.Addr().String()
	} else {
		ready <- ss[len(ss)-1]
	}
}

func handleConn(conn *net.TCPConn) {
	defer func() { _ = conn.Close() }()
	buf := make([]byte, maxBufferSize)
	if err := conn.SetReadBuffer(maxBufferSize); err != nil {
		panic(err)
	}
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	if err := conn.SetWriteBuffer(n); err != nil {
		panic(err)
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	if _, err := conn.Write(buf); err != nil {
		panic(err)
	}
}
