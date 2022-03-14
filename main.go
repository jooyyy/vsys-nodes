package main

import (
	"github.com/jooyyy/vsys-nodes/server"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

const (
	Version = "0.0.1"
)

func main() {
	c := cli.NewCLI("vsys-nodes", Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"api": server.CommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
