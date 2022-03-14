package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jooyyy/vsys-nodes/config"
	"github.com/mitchellh/cli"
	"github.com/panjf2000/ants"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Command struct{}

func CommandFactory() (cli.Command, error) {
	return new(Command), nil
}

type Service struct {
	mode       string
	httpServer *http.Server
}

func NewService(args ...string) *Service {
	handler := gin.Default()

	addr := fmt.Sprintf(":%d", config.GetConfig().Port)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	service := &Service{
		httpServer: server,
	}

	service.initRouter(handler)

	return service
}

func (s *Service) Run() int {

	startRestService := func() {
		fmt.Println("Start rest api server", s.httpServer.Addr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			fmt.Errorf("RestServer.Run %s", err)
		}
		fmt.Println("rest server shutdown.")
	}

	waitToStop := func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
	}

	stopAllService := func() {
		s.httpServer.Shutdown(context.Background())
		taskPool.Release()
		fmt.Println("all service has stopped, exit.")
	}

	go startRestService()

	taskPool = initTaskPool()
	nodes = initNodes()

	startNodesHealthCheck()

	waitToStop()
	stopAllService()

	return 1
}

var taskPool *ants.PoolWithFunc

var nodes []Node

func (c *Command) Run(args []string) int {
	return NewService(args...).Run()
}

func (c *Command) Help() string {
	return help
}

func (c *Command) Synopsis() string {
	return synopsis
}

const synopsis = "restful api service tips"
const help = `
Usage: api xxxx
`

func initTaskPool() *ants.PoolWithFunc {
	p, _ := ants.NewPoolWithFunc(20, func(i interface{}) {
		if t, ok := i.(NodeHealthTask); ok {
			t.Do()
		}else {
			fmt.Println("unknown task type")
		}
	})
	return p
}

func initNodes() []Node {
	var out []Node
	for _, one := range config.GetConfig().Mainnet {
		out = append(out, Node{
			Url:          one,
			Network:      mainnet,
		})
	}
	for _, one := range config.GetConfig().Testnet {
		out = append(out, Node{
			Url:          one,
			Network: 	  testnet,
		})
	}
	return out
}
