package main

import (
	"log"
	"strings"

	"go.uber.org/dig"
)

type A struct {
	*C
}

const indent = 8

func (a A) callA(depth int) {
	prefix := strings.Repeat(" ", depth*indent)
	log.Println(prefix, "A in")
	log.Println(prefix, "-> calling C")
	a.callC(depth + 1)

	log.Println(prefix, "A out")
}

type B struct {
	*A
	*C
}

func (b B) callB(depth int) {
	prefix := strings.Repeat(" ", depth*indent)
	log.Println(prefix, "B in")
	log.Println(prefix, "-> calling A")
	b.callA(depth + 1)

	log.Println(prefix, "B in")
	log.Println(prefix, "-> calling C")
	b.callC(depth + 1)

	log.Println(prefix, "B out")
}

type C struct{}

func (c C) callC(depth int) {
	prefix := strings.Repeat(" ", depth*indent)
	log.Println(prefix, "C in")
	log.Println(prefix, "C out")
}

type Tot struct {
	*A
	*B
}

func main() {
	container := dig.New()

	if err := container.Provide(func(c *C) *A { return &A{C: c} }); err != nil {
		log.Fatalln(err)
	}
	if err := container.Provide(func(a *A, c *C) *B { return &B{A: a, C: c} }); err != nil {
		log.Fatalln(err)
	}
	if err := container.Provide(func() *C { return &C{} }); err != nil {
		log.Fatalln(err)
	}
	if err := container.Provide(func(a *A, b *B) *Tot { return &Tot{A: a, B: b} }); err != nil {
		log.Fatalln(err)
	}

	if err := container.Invoke(func(t *Tot) {
		t.callA(0)
		t.callB(0)
	}); err != nil {
		log.Fatalln(err)
	}
}
