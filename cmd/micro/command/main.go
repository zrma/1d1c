package main

import (
	"fmt"
	"strings"

	"context"
	"github.com/micro/go-micro"

	proto "github.com/micro/go-micro/agent/proto"
)

// Command implements CommandHelper
type Command struct{}

// Help returns the command usage
func (c *Command) Help(ctx context.Context, req *proto.HelpRequest, rsp *proto.HelpResponse) error {
	rsp.Usage = "command"
	rsp.Description = "This is an example bot command as a micro service"
	return nil
}

// Exec executes the command
func (c *Command) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) error {
	rsp.Result = []byte(strings.Join(req.Args, " "))
	// rsp.Error could be set to return an error instead
	// the function error would only be used for service level issues
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.bot.command"),
	)

	service.Init()

	if err := proto.RegisterCommandHandler(service.Server(), new(Command)); err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
