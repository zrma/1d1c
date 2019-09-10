package main

import (
	"github.com/zrma/1d1c/cmd/interface/vehicle"
	"log"
)

type car struct {
	dist int
}

var _ vehicle.Vehicle = (*car)(nil)

func (c *car) Dist() int {
	return c.dist
}

func (c *car) Run() {
	c.dist++
}

func main() {
	var v vehicle.Vehicle = &car{}
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
		log.Println("v is not Dist() interface")
	}
	log.Println(d)
	log.Println(d.Dist())
}
