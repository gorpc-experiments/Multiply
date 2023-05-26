package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorpc-experiments/ServiceCore"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	spew.Dump(args, reply)
	return nil
}

func main() {
	arith := new(Arith)
	err := rpc.Register(arith)

	client, err := ServiceCore.NewGalaxyClient()

	if err != nil {
		log.Println(err.Error())
		return
	}

	client.RegisterToGalaxy(arith)

	rpc.HandleHTTP()
	port := ServiceCore.GetRPCPort()

	println("Multiply is running on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println(err.Error())
	}
}
