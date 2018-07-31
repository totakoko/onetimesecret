package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/common"
)

type HTTPServer struct {
	Store common.Store

	router *gin.Engine
}

func (s *HTTPServer) Init() {

	gin.SetMode(gin.ReleaseMode)
	s.router = gin.New()
	// s.router.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true), gin.Recovery())

	s.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	s.router.POST("/secrets", s.SecretPost)
	s.router.GET("/secrets/:id", s.SecretGet)
}

func (s *HTTPServer) Run(addr string) error {
	log.Error().Msgf("Listening on %s", addr)
	return s.router.Run(addr)
}
