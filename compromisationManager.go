package compromiseAnalyser

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var (
	hash = ""
)

func AsyncRiskEvaluator(system string, verbose bool, port int) {
	go asyncEvaluate(system, verbose, port)
}

func asyncEvaluate(system string, verbose bool, port int) {
	setupRPC(port)
	for {
		hashResult, err := MakeAllCheck(system)
		if verbose {
			fmt.Printf("Error while executing check: %s ", err.Error())
		}
		hash = hex.EncodeToString(hashResult)
	}
}

func setupRPC(port int) {
	cm := new(CompManager)
	rpc.Register(cm)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+string(port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

type Args struct {
}

type Response struct {
	Hash string
}

type CompManager int

func (t *CompManager) Status(args *Args, reply *Response) error {
	reply.Hash = hash
	return nil
}
