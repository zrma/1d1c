package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-test/deep"
)

// PrintTest 함수는 인자로 전달하는 함수 내부에서 표준 출력으로 출력한 문자열과 두 번째 인자로 전달한 문자열을 비교한다.
// 만약 그 내용이 다르면 오류를 반환한다.
func PrintTest(f func(), expected []string) (err error) {
	r, w, _ := os.Pipe()
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
		_ = r.Close()
	}()
	os.Stdout = w
	go func() {
		defer func() {
			_ = w.Close()
		}()
		f()
	}()
	var stdout []byte
	stdout, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	const lineFeed = "\n"
	got := strings.Split(string(stdout), lineFeed)
	if got[len(got)-1] == "" {
		got = got[:len(got)-1]
	}
	diff := deep.Equal(got, expected)
	if diff != nil {
		return errors.New(strings.Join(diff, lineFeed))
	}
	return nil
}
