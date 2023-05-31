package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gorpc-experiments/ServiceCore"
)

type Args struct {
	A, B int
}

type Arith struct {
	ServiceCore.CoreHealth
}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	spew.Dump(args, reply)
	return nil
}

func main() {
	ServiceCore.SetupLogging()

	arith := new(Arith)

	ServiceCore.PublishMicroService(arith, true)
}