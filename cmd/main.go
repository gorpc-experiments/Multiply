package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gorpc-experiments/GalaxyClient"
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

	client, err := GalaxyClient.NewGalaxyClient()

	if err != nil {
		log.Println(err.Error())
		return
	}

	client.RegisterToGalaxy(arith)

	rpc.HandleHTTP()

	err = http.ListenAndServe(":3456", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
