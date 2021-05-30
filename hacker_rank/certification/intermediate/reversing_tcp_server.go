package intermediate

import (
	"context"
	"io"
	"net"
)

const (
	maxBufferSize = 1024
	network       = "tcp"
)

type TCPAcceptor interface {
	AcceptTCP() (*net.TCPConn, error)
}

type Logger interface {
	Error(args ...interface{})
}

func TCPServer(ctx context.Context, acceptor TCPAcceptor, logger Logger) chan interface{} {
	waitCh := make(chan interface{})
	go func() {
		defer close(waitCh)
		for ctx.Err() == nil {
			conn, err := acceptor.AcceptTCP()
			if ctx.Err() != nil && ctx.Err() == context.Canceled {
				return
			}
			if err != nil {
				logger.Error("TCPServer", err)
				return
			}
			go handleConn(conn)
		}
	}()
	return waitCh
}

type Buffer interface {
	SetReadBuffer(int) error
	SetWriteBuffer(int) error
	io.ReadWriter
	io.Closer
}

func handleConn(conn Buffer) {
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
