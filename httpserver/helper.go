package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/common/errors"
)

func sendErrorResponse(c *gin.Context, err error) {
	log.Error().Err(err)
	c.JSON(getStatusFromError(err), gin.H{
		"message": err.Error(),
	})
}

func getStatusFromError(err error) int {
	switch err.(type) {
	case *errors.AppError:
		log.Error().Fields(map[string]interface{}{
			"error": err.Error(),
		}).Msg("Service error")
		return err.(*errors.AppError).HTTPCode
	default:
		log.Error().Fields(map[string]interface{}{
			"error": err.Error(),
		}).Msg("Unknown service error")
		return 500
	}
}
