package main

import (
	"os"

	"gx/ipfs/QmP9vZfc5WSjfGTXmwX2EcicMFzmZ6fXn7HTdKYat6ccmH/go-ipfs-cmds/examples/adder"

	"gx/ipfs/QmP9vZfc5WSjfGTXmwX2EcicMFzmZ6fXn7HTdKYat6ccmH/go-ipfs-cmds"
	"gx/ipfs/QmP9vZfc5WSjfGTXmwX2EcicMFzmZ6fXn7HTdKYat6ccmH/go-ipfs-cmds/cli"
)

func main() {
	// parse the command path, arguments and options from the command line
	req, cmd, _, err := cli.Parse(os.Args[1:], os.Stdin, adder.RootCmd)
	if err != nil {
		panic(err)
	}

	req.SetOption("encoding", cmds.Text)

	// create an emitter
	re, retCh := cli.NewResponseEmitter(os.Stdout, os.Stderr, cmd.Encoders["Text"], req)

	if pr, ok := cmd.PostRun[cmds.CLI]; ok {
		re = pr(req, re)
	}

	wait := make(chan struct{})
	// call command in background
	go func() {
		defer close(wait)

		err = adder.RootCmd.Call(req, re)
		if err != nil {
			panic(err)
		}
	}()

	// wait until command has returned and exit
	ret := <-retCh
	<-wait

	os.Exit(ret)
}
