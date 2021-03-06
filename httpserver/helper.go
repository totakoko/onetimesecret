package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/totakoko/onetimesecret/common/errors"
	"gitlab.com/totakoko/onetimesecret/helpers"
)

func sendErrorResponse(c *gin.Context, err error) {
	log.Error().Fields(map[string]interface{}{
		"error": err.Error(),
	}).Msg("Service error")
	c.JSON(getStatusFromError(err), gin.H{
		"message": err.Error(),
	})
}
func (s *HTTPServer) sendErrorPage(c *gin.Context, err error) {
	log.Error().Fields(map[string]interface{}{
		"error": err.Error(),
	}).Msg("Service error")
	c.Status(getStatusFromError(err))
	helpers.LogOnError(s.templatesCache["error"].Execute(c.Writer, map[string]interface{}{
		"error": err.Error(),
	}))
}

func getStatusFromError(err error) int {
	switch err.(type) {
	case *errors.AppError:
		return err.(*errors.AppError).HTTPCode
	default:
		return 500
	}
}
