package middleware

import (
	"fmt"
	"net/http"

	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messgae": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Print(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messgae": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
