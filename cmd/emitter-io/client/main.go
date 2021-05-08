package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/emitter-io/go"
)

func main() {
	// Create the options with default values
	o := emitter.NewClientOptions()
	o.AddBroker("tcp://localhost:8080")

	// Set the message handler
	o.SetOnMessageHandler(func(client emitter.Emitter, msg emitter.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	})

	// Create a new emitter client and connect to the broker
	c := emitter.NewClient(o)
	sToken := c.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		panic("Error on Client.Connect(): " + sToken.Error().Error())
	}

	// Subscribe to the presence demo channel
	c.Subscribe("I4dcTh0TFQzJekBgTXrigIf1wooI2-EC", "abc/1/")

	pid := strconv.Itoa(os.Getpid())
	fmt.Println("I'am " + pid)

	// Publish to the channel
	c.Publish("I4dcTh0TFQzJekBgTXrigIf1wooI2-EC", "abc/1/", pid+": hello")
	time.Sleep(time.Second * 1)

	c.Publish("I4dcTh0TFQzJekBgTXrigIf1wooI2-EC", "abc/1/", pid+": greeting")
	time.Sleep(time.Second * 1)

	c.Publish("I4dcTh0TFQzJekBgTXrigIf1wooI2-EC", "abc/1/", pid+": bye")
	time.Sleep(time.Second * 1)

	// Ask for presence
	r := emitter.NewPresenceRequest()
	r.Key = "I4dcTh0TFQzJekBgTXrigIf1wooI2-EC"
	r.Channel = "abc/1/"
	c.Presence(r)

	time.Sleep(time.Second * 3)
}
