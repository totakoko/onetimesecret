package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/common/errors"
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
	s.templatesCache["error"].Execute(c.Writer, map[string]interface{}{
		"error": err.Error(),
	})
}

func getStatusFromError(err error) int {
	switch err.(type) {
	case *errors.AppError:
		return err.(*errors.AppError).HTTPCode
	default:
		return 500
	}
}
