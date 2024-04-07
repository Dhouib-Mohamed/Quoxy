package validation

import (
	"github.com/gin-gonic/gin"
)

func ValidateBodyMiddleware(fields []field) gin.HandlerFunc {
	return func(context *gin.Context) {
		// get body

	}
}
