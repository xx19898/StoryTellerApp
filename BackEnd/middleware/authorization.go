// TODO: implement and test authorization
package middleware

import (
	"StoryTellerAppBackend/helpers"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		headersWithAuthorizationKey, authHeadersFound := c.Request.Header["Authorization"]

		if !authHeadersFound || len(headersWithAuthorizationKey) == 0 {
			c.JSON(403, gin.H{
				"message": "No access token found",
			})
			return
		}

		secret, ok := helpers.GetEnv("JWT_SECRET")

		authToken := headersWithAuthorizationKey[0]

		if !ok {
			c.JSON(500, gin.H{
				"message": "Internal error",
			})
			return
		}

		_, roles, errRolesExtract := ExtractRolesAndUsername(authToken, secret)

		if errRolesExtract != nil {
			c.JSON(500, gin.H{
				"message": "Internal error",
			})
			return
		}

		c.Set("ROLES", roles)

		c.Next()
	}
}
