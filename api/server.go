package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nportas/alfonsina-bot/poemas"
)

type GinServer struct {
	poemario *poemas.Poemario
}

func NewGinServer(poemario *poemas.Poemario) *GinServer {
	return &GinServer{poemario}
}

func (server *GinServer) Start() {

	router := gin.Default()

	router.GET("/generarPoesia/:palabra", server.generarPoesia)

	router.Run()
}

func (server *GinServer) generarPoesia(c *gin.Context) {

	palabra := c.Param("palabra")
	c.JSON(http.StatusOK, server.poemario.GenerarPoesiaAPartirDe(palabra))
}
