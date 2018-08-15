package httpserver

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/common"
)

// changes to these paths must be correlated with changes in gulpfile.js
const publicPath = ".build/public"
const templatesPath = ".build/templates"

type HTTPServer struct {
	PublicURL string
	Store     common.Store

	router         *gin.Engine
	assetsCache    map[string][]byte
	templatesCache map[string]*template.Template
}

func (s *HTTPServer) Init() error {
	s.assetsCache = make(map[string][]byte)
	s.templatesCache = make(map[string]*template.Template)

	gin.SetMode(gin.ReleaseMode)
	s.router = gin.New()

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

	if err := s.loadTemplates(); err != nil {
		return err
	}
	if err := s.loadStaticAssets(); err != nil {
		return err
	}

	s.router.GET("/test/secret", s.TestSecret)
	s.router.GET("/test/link", s.TestCreateSecret)
	s.router.GET("/", s.DisplayHomePage)
	s.router.GET("/about", s.DisplayAboutPage)
	s.router.POST("/secrets", s.CreateSecret)
	s.router.GET("/secrets/:id", s.GetSecret)
	s.router.POST("/api/secrets", s.APICreateSecret)
	s.router.GET("/api/secrets/:id", s.APIGetSecret)
	return nil
}

func (s *HTTPServer) Run(addr string) error {
	log.Warn().Msgf("Listening on %s", addr)
	return s.router.Run(addr)
}

func (s *HTTPServer) TestSecret(c *gin.Context) {
	c.Status(http.StatusOK)
	s.templatesCache["view-secret"].Execute(c.Writer, map[string]interface{}{
		"secret": "mysecret",
	})
}

func (s *HTTPServer) TestCreateSecret(c *gin.Context) {
	c.Status(http.StatusOK)
	s.templatesCache["view-secret-link"].Execute(c.Writer, map[string]interface{}{
		"secretURL": "http://localhost:5000/test/secret",
	})
}
