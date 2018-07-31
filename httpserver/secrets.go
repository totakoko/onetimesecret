package httpserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) SecretPost(c *gin.Context) {
	expiration, err := strconv.ParseUint(c.PostForm("expiration"), 10, 64)
	secretKey, err := s.Store.StoreSecret(c.PostForm("secret"), time.Duration(expiration)*time.Second)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, secretKey)
}

func (s *HTTPServer) SecretGet(c *gin.Context) {
	secret, err := s.Store.GetSecret(c.Param("secretId"))
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, secret)
}
