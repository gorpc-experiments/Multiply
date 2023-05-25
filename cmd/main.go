package main

import (
	"fmt"
	"github.com/AliceDiNunno/KubernetesUtil"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorpc-experiments/GalaxyClient"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
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

func getPort() int {
	port := 0
	if KubernetesUtil.IsRunningInKubernetes() {
		port = KubernetesUtil.GetInternalServicePort()
	}
	if port == 0 {
		env_port := os.Getenv("PORT")
		if env_port == "" {
			log.Fatalln("PORT env variable isn't set")
		}
		envport, err := strconv.Atoi(env_port)
		if err != nil {
			log.Fatalln(err.Error())
		}
		port = envport
	}

	return port
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
	port := getPort()

	println("Multiply is running on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println(err.Error())
	}
}
