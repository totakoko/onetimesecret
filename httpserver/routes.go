package httpserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/totakoko/onetimesecret/common/errors"
	"gitlab.com/totakoko/onetimesecret/helpers"
)

func (s *HTTPServer) DisplayHomePage(c *gin.Context) {
	helpers.LogOnError(s.templatesCache["create-secret"].Execute(c.Writer, nil))
}
func (s *HTTPServer) DisplayOfflinePage(c *gin.Context) {
	helpers.LogOnError(s.templatesCache["_offline"].Execute(c.Writer, nil))
}

func (s *HTTPServer) DisplayAboutPage(c *gin.Context) {
	helpers.LogOnError(s.templatesCache["about"].Execute(c.Writer, nil))
}

func (s *HTTPServer) CreateSecret(c *gin.Context) {
	expiration, err := strconv.ParseUint(c.PostForm("expiration"), 10, 64)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	secretKey, err := s.Store.StoreSecret(c.PostForm("secret"), time.Duration(expiration)*time.Second)
	if err != nil {
		s.sendErrorPage(c, err)
		return
	}
	c.Status(http.StatusCreated)
	helpers.LogOnError(s.templatesCache["view-secret-link"].Execute(c.Writer, map[string]interface{}{
		"secretURL":  s.PublicURL + "secrets/" + secretKey,
		"expiration": c.PostForm("expiration") + " seconds",
	}))
}

func (s *HTTPServer) GetSecret(c *gin.Context) {
	c.Status(http.StatusOK)
	helpers.LogOnError(s.templatesCache["get-secret"].Execute(c.Writer, map[string]interface{}{
		"secretViewURL": "/secrets/" + c.Param("id") + "/view",
	}))
}
func (s *HTTPServer) GetSecretContent(c *gin.Context) {
	secret, err := s.Store.GetSecret(c.Param("id"))
	switch err.(type) {
	case nil:
		c.Status(http.StatusOK)
		helpers.LogOnError(s.templatesCache["view-secret"].Execute(c.Writer, map[string]interface{}{
			"secret": secret,
		}))
	case *errors.AppError:
		c.Status(http.StatusNotFound)
		helpers.LogOnError(s.templatesCache["view-secret"].Execute(c.Writer, map[string]interface{}{
			"secret": "Unknown secret",
		}))
	default:
		sendErrorResponse(c, err)
	}
}
