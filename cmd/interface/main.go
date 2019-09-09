package main

import (
	"log"
)

type vehicle interface {
	Run()
}

type car struct {
	dist int
}

func (c *car) Dist() int {
	return c.dist
}

func (c *car) Run() {
	c.dist++
}

func main() {
	var v vehicle = &car{}
	v.Run()
	v.Run()
	v.Run()

	g, ok := v.(interface{ Get() int })
	if !ok {
		log.Println("v is not Get() interface")
	}
	log.Println(g)

	d, ok := v.(interface{ Dist() int })
	if !ok {
		log.Println("v is not Get() interface")
	}
	log.Println(d)
	log.Println(d.Dist())
}
