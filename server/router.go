package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jooyyy/vsys-nodes/server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "net/http/pprof"
)

func (s *Service) initRouter(r gin.IRouter) {
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/api/node/:network", s.GetNode)
	r.GET("/api/node/all", s.GetAllNodes)
}
