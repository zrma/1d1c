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
			addrChan := make(chan string)
			ctx, cancel := context.WithCancel(context.Background())
			go TCPServer(ctx, addrChan)
			addr := <-addrChan
			got, err := tcpClient(addr, tt.given)
			if err != nil {
				t.Log(err.Error())
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
			cancel()
		})
	}
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
