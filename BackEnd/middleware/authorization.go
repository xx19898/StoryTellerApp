package middleware

import (
	"StoryTellerAppBackend/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserInfoExtractionMiddleware() gin.HandlerFunc {
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

		username, roles, errRolesExtract := ExtractUserInfo(authToken, secret)

		if errRolesExtract != nil {
			c.JSON(500, gin.H{
				"message": "Internal error",
			})
			return
		}

		c.Set("ROLES", roles)
		c.Set("LOGGED_USER_NAME", username)
		fmt.Print(username)
		c.Next()
	}
}

func AuthorizationMiddleware(compareRoles func([]string, []string) bool, neededRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		grantedRoles := c.GetStringSlice("ROLES")
		result := compareRoles(grantedRoles, neededRoles)
		if !result {
			c.JSON(403, gin.H{
				"message": "You are not authorized to access this resource",
			})
		}
		c.Next()
	}
}

func CompareRoles(rolesFound []string, rolesNeeded []string) bool {
	var roleIsFound bool

	for _, roleToFind := range rolesNeeded {
		roleIsFound = false
		for _, roleFound := range rolesFound {
			if roleFound == roleToFind {
				roleIsFound = true
				break
			}
		}
		if !roleIsFound {
			break
		}
	}

	return roleIsFound
}
