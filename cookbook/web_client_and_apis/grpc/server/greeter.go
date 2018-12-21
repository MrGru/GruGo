package main

import (
	"context"
	"fmt"

	"github.com/MrGru/GruGo/cookbook/web_client_and_apis/grpc/greeter"
)

type Greeter struct {
	Exclaim bool
}

func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}
	return &greeter.GreetResponse{Response: msg}, nil
}
