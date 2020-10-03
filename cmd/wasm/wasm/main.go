// +build js,wasm

package main

//go:generate env GOOS=js GOARCH=wasm go build -o ../server/static/hello.wasm main.go

func main() {
	println("Hello, wasm!")
}
