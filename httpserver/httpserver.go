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
	PublicURL string
	Store     common.Store

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

	s.router.Static("/assets", ".build/assets")

	s.router.GET("/", s.DisplayHomePage)
	s.router.GET("/about", s.DisplayAboutPage)
	s.router.POST("/secrets", s.CreateSecret)
	s.router.GET("/secrets/:id", s.GetSecret)
	s.router.POST("/api/secrets", s.APICreateSecret)
	s.router.GET("/api/secrets/:id", s.APIGetSecret)
}

func (s *HTTPServer) Run(addr string) error {
	log.Warn().Msgf("Listening on %s", addr)
	return s.router.Run(addr)
}
