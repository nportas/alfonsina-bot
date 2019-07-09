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

	router.GET("/generarPoesiaAPartirDe/:palabra/:cantidadMinVersos/:cantidadMaxVersos", server.generarPoesiaAPartirDe)
	router.GET("/generarPoesia/:cantidadMinVersos/:cantidadMaxVersos", server.generarPoesia)

	router.Run()
}

func (server *GinServer) generarPoesiaAPartirDe(c *gin.Context) {
	palabra := c.Param("palabra")
	cantidadMinVersos := c.Param("cantidadMinVersos")
	cantidadMaxVersos := c.Param("cantidadMaxVersos")
	c.JSON(http.StatusOK, server.poemario.GenerarPoesiaAPartirDe(palabra, cantidadMinVersos, cantidadMaxVersos))
}

func (server *GinServer) generarPoesia(c *gin.Context) {
	cantidadMinVersos := c.Param("cantidadMinVersos")
	cantidadMaxVersos := c.Param("cantidadMaxVersos")
	c.JSON(http.StatusOK, server.poemario.GenerarPoesia(cantidadMinVersos, cantidadMaxVersos))
}
