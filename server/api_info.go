package server

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	mainnet = "mainnet"
	testnet = "testnet"
)

type Node struct {
	Url 	string `json:"url"`
	Height 	int64 `json:"height"`
	Network string `json:"network"`
	Weight int `json:"weight,omitempty"`
	Point int `json:"point,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	TimeSinceNow string `json:"time_since_now,omitempty"`
	Note string `json:"note,omitempty"`
}

// GetNode
// @Summary Get node by network type
// @Tags Nodes
// @Accept  json
// @Produce json
// @Params token query string false "api token"
// @Success 200 {object} string "return an active node url"
// @Failure 400 {object} RespJsonObj ""
// @Router /api/node/:network [get]
func (s *Service) GetNode(c *gin.Context) {
	network := c.Param("network")
	switch strings.ToLower(network) {
	case mainnet:
		return
	case testnet:
		return
	default:
		respErr(c, "unknown network type")
		break
	}
}

// GetAllNodes
// @Summary Get all nodes information
// @Tags Nodes
// @Accept  json
// @Produce json
// @Params token query string false "api token"
// @Success 200 {object} []Node "return an active node url"
// @Failure 400 {object} RespJsonObj ""
// @Router /api/node/all [get]
func (s *Service) GetAllNodes(c *gin.Context) {
	var out = []Node{}
	for _, one := range nodes {
		out = append(out, Node{
			Url:          one.Url,
			Height:       one.Height,
			Network:      one.Network,
			Timestamp:    one.Timestamp,
			TimeSinceNow: time.Since(time.Unix(one.Timestamp, 0)).String(),
			Note:         one.Note,
		})
	}
	respJson(c, out)
}
