package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

func HandleErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case apperr.AppError:
				log.Printf("ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": fmt.Sprintf("%d: %s", e.Code(), e.Msg()),
				})
			default:
				log.Printf("FATAL: %+v", e)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Fatal",
				})
			}

		}
	}
}
