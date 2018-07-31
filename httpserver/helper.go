package httpserver

import (
	"fmt"

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
	switch t := err.(type) {
	default:
		fmt.Println("not a model missing error")
		return 500
	case *errors.AppError:
		fmt.Println("AppError", t)
		return err.(*errors.AppError).HTTPCode
	}
}
