package intermediate

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPServer(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		given []string
		want  []string
	}{
		{
			[]string{"abc", "def", "ghi"},
			[]string{"cba", "fed", "ihg"},
		},
		{
			[]string{"eno", "owt", "eerht", "ruof"},
			[]string{"one", "two", "three", "four"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			listener := newTCPListener(t)
			addr := parseAddr(listener.Addr().String())

			logger := mockLogger{}

			waitCh := TCPServer(ctx, listener, &logger)
			got, err := tcpClient(addr, tt.given)
			if err != nil {
				t.Log(err.Error())
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)

			cancel()
			assert.NoError(t, listener.Close())
			<-waitCh

			assert.Empty(t, logger.got)
		})
	}
}

func newTCPListener(t *testing.T) *net.TCPListener {
	const endpoint = ""
	addr, err := net.ResolveTCPAddr(network, endpoint)
	assert.NoError(t, err)

	listener, err := net.ListenTCP(network, addr)
	assert.NoError(t, err)

	return listener
}

func parseAddr(original string) string {
	parsed := strings.Split(original, ":")
	if len(parsed) < 2 {
		return original
	} else {
		return parsed[len(parsed)-1]
	}
}

func TestParseAddr(t *testing.T) {
	for i, tt := range []struct {
		given string
		want  string
	}{
		{"[::]:1234", "1234"},
		{":1234", "1234"},
		{"", ""},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := parseAddr(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTCPServerAcceptFailure(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const wantErrMsg = "this is sparta"
	logger := mockLogger{}
	waitCh := TCPServer(ctx, mockAcceptor{wantErrMsg}, &logger)
	<-waitCh

	assert.Len(t, logger.got, 1)
	got := logger.got[0]

	assert.Len(t, got, 2)
	assert.Equal(t, "TCPServer", got[0])

	err, ok := got[1].(error)
	assert.True(t, ok)
	assert.EqualError(t, err, wantErrMsg)
}

type mockLogger struct {
	got [][]interface{}
}

func (m *mockLogger) Error(arg ...interface{}) {
	m.got = append(m.got, arg)
}

type mockAcceptor struct {
	wantErrMsg string
}

func (m mockAcceptor) AcceptTCP() (*net.TCPConn, error) {
	return nil, fmt.Errorf(m.wantErrMsg)
}

func TestHandleConn(t *testing.T) {
	assert.NotPanics(t, func() {
		const (
			given = "abc"
			want  = "cba"
		)

		b := newBufferImpl(t, given, want)
		handleConn(b)
	})
}

func TestHandleConnError(t *testing.T) {
	const (
		given = "abc"
		want  = "cba"
	)

	t.Run("SetReadBuffer", func(t *testing.T) {
		const wantErrMsg = "SetReadBuffer"
		assert.PanicsWithError(t, wantErrMsg, func() {
			b := newBufferImpl(t, given, want)
			b.setReadBuffer = func(n int) error {
				assert.Equal(t, maxBufferSize, n)
				return fmt.Errorf(wantErrMsg)
			}
			handleConn(b)
		})
	})

	t.Run("SetWriteBuffer", func(t *testing.T) {
		const wantErrMsg = "SetWriteBuffer"
		assert.PanicsWithError(t, wantErrMsg, func() {
			b := newBufferImpl(t, given, want)
			b.setWriteBuffer = func(n int) error {
				assert.Equal(t, len(want), n)
				return fmt.Errorf(wantErrMsg)
			}
			handleConn(b)
		})
	})

	t.Run("Read", func(t *testing.T) {
		const wantErrMsg = "read 실패"
		assert.PanicsWithError(t, wantErrMsg, func() {
			b := newBufferImpl(t, given, want)
			b.read = func(p []byte) (n int, err error) {
				assert.NotNil(t, p)
				assert.Len(t, p, maxBufferSize)
				return 0, fmt.Errorf(wantErrMsg)
			}
			handleConn(b)
		})
	})

	t.Run("Write", func(t *testing.T) {
		const wantErrMsg = "write 실패"
		assert.PanicsWithError(t, wantErrMsg, func() {
			b := newBufferImpl(t, given, want)
			b.write = func(p []byte) (n int, err error) {
				assert.NotNil(t, p)
				assert.Len(t, p, len(want))
				for i, c := range want {
					assert.EqualValues(t, c, p[i])
				}
				return len(given), fmt.Errorf(wantErrMsg)
			}
			handleConn(b)
		})
	})
}

func newBufferImpl(t *testing.T, given, want string) bufferImpl {
	return bufferImpl{
		setReadBuffer: func(n int) error {
			assert.Equal(t, maxBufferSize, n)
			return nil
		},
		setWriteBuffer: func(n int) error {
			assert.Equal(t, len(want), n)
			return nil
		},
		read: func(p []byte) (n int, err error) {
			assert.NotNil(t, p)
			assert.Len(t, p, maxBufferSize)
			for i, c := range given {
				p[i] = byte(c)
			}
			return len(given), nil
		},
		write: func(p []byte) (n int, err error) {
			assert.NotNil(t, p)
			assert.Len(t, p, len(want))
			for i, c := range want {
				assert.EqualValues(t, c, p[i])
			}
			return len(given), nil
		},
	}
}

type bufferImpl struct {
	setReadBuffer  func(n int) error
	setWriteBuffer func(n int) error
	read           func(p []byte) (n int, err error)
	write          func(p []byte) (n int, err error)
}

func (b bufferImpl) SetReadBuffer(i int) error {
	if b.setReadBuffer != nil {
		return b.setReadBuffer(i)
	}
	panic("implement me")
}

func (b bufferImpl) SetWriteBuffer(i int) error {
	if b.setWriteBuffer != nil {
		return b.setWriteBuffer(i)
	}
	panic("implement me")
}

func (b bufferImpl) Read(p []byte) (n int, err error) {
	if b.read != nil {
		return b.read(p)
	}
	panic("implement me")
}

func (b bufferImpl) Write(p []byte) (n int, err error) {
	if b.write != nil {
		return b.write(p)
	}
	panic("implement me")
}

func (b bufferImpl) Close() error {
	return nil
}

func tcpClient(addr string, messages []string) ([]string, error) {
	if !strings.Contains(addr, ":") {
		addr = fmt.Sprintf(":%s", addr)
	}

	tcpAddr, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		return []string{}, err
	}

	var reversed []string
	for _, msg := range messages {

		conn, err := net.DialTCP(network, nil, tcpAddr)
		if err != nil {
			return []string{}, err
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			return []string{}, err
		}

		reply := make([]byte, maxBufferSize)

		n, err := conn.Read(reply)
		if err != nil {
			return []string{}, err
		}

		reversed = append(reversed, string(reply[:n]))
		_ = conn.Close()
	}

	return reversed, nil
}
