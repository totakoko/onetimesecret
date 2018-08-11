package httpserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) APICreateSecret(c *gin.Context) {
	expiration, err := strconv.ParseUint(c.PostForm("expiration"), 10, 64)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	secretKey, err := s.Store.StoreSecret(c.PostForm("secret"), time.Duration(expiration)*time.Second)
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, secretKey)
}

func (s *HTTPServer) APIGetSecret(c *gin.Context) {
	secret, err := s.Store.GetSecret(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, secret)
}
