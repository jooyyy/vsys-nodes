package server

import (
	"encoding/json"
	"fmt"
	"github.com/jooyyy/vsys-sdk-go/vsys"
	"github.com/robfig/cron/v3"
)

type NodeHealthTask struct {
	node 	*Node
	handler func(state nodeStateResp)
	errHandler func(msg string)
}

type nodeStateResp struct {
	BlockchainHeight int64 `json:"blockchainHeight"`
	StateHeight int64 `json:"stateHeight"`
	UpdatedTimestamp int64 `json:"updatedTimestamp"`
	UpdateDate string `json:"updateDate"`
}

func (t *NodeHealthTask) Do() {
	content, err := vsys.UrlGetContent(t.node.Url + "/node/status")
	if err != nil {
		t.node.Note = err.Error()
		t.errHandler(err.Error())
		return
	}

	var resp nodeStateResp
	if err = json.Unmarshal(content, &resp);err != nil {
		t.node.Note = err.Error()
		t.errHandler(err.Error())
		return
	}

	t.node.Timestamp = resp.UpdatedTimestamp/1e9
	t.node.Height = resp.StateHeight
	t.node.Note = ""
}


func startNodesHealthCheck() {
	c := cron.New()
	c.AddFunc("@every 10m", func() {
		for i := range nodes {
			_ = taskPool.Invoke(NodeHealthTask{
				node:       &nodes[i],
				handler: func(state nodeStateResp) {
					fmt.Println(state)
				},
				errHandler: func(msg string) {
					fmt.Println(msg)
				},
			})
		}
	})
	c.Start()
}
