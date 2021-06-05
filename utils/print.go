package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// GetPrinted 함수는
// 첫 번째 인자로 전달한 함수(f)의 내부에서 표준 출력으로 출력한 문자열을 반환한다.
func GetPrinted(f func()) ([]string, error) {
	r, w, _ := os.Pipe()
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
		_ = r.Close()
	}()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			_ = w.Close()
			wg.Done()
		}()
		f()
	}()
	wg.Wait()
	stdout, err := readAll(r)
	if err != nil {
		return nil, err
	}
	const lineFeed = "\n"
	got := strings.Split(string(stdout), lineFeed)
	if got[len(got)-1] == "" {
		got = got[:len(got)-1]
	}
	return got, nil
}

var readAll = ioutil.ReadAll
