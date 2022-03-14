package server

import (
	"github.com/gin-gonic/gin"
)

const (
	OK                  = 0
	Redirect301         = 301
	Redirect302         = 302
	BadRequest          = 400
	ErrorRequireLogin   = 403
	InternalServerError = 500
)

var statusText = map[int]string{
	OK:                  "success",
	BadRequest:          "bad request params",
	InternalServerError: "server internal error",
}

type RespJsonObj struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func respErr(c *gin.Context, msg string) {
	c.JSON(BadRequest, RespJsonObj{
		Code: BadRequest,
		Msg:  msg,
	})
}

func respJson(c *gin.Context, data interface{}) {
	c.JSON(OK, RespJsonObj{
		Code: OK,
		Msg:  statusText[OK],
		Data: data,
	})
}

func respAuthFail(c *gin.Context, msg string) {
	c.JSON(ErrorRequireLogin, RespJsonObj{
		Code: ErrorRequireLogin,
		Msg:  msg,
	})
}
