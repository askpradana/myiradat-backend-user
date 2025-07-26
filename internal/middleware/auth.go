package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"myiradat-backend-auth/internal/response"
)

type ServiceClaim struct {
	ServiceName string `json:"serviceName"`
	RoleName    string `json:"roleName"`
	ServiceCode string `json:"serviceCode"`
}

type CustomClaims struct {
	Email    string         `json:"email"`
	Services []ServiceClaim `json:"services"`
	jwt.RegisteredClaims
}

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, "invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			response.ServerError(c, "JWT secret not set")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, "invalid or expired token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			response.Error(c, "invalid token claims")
			c.Abort()
			return
		}

		var dashRole string
		for _, svc := range claims.Services {
			if svc.ServiceCode == "DASHBOARD" {
				dashRole = svc.RoleName
				break
			}
		}

		if dashRole == "" {
			response.Error(c, "no access to service DASH")
			c.Abort()
			return
		}

		roleAllowed := false
		for _, role := range allowedRoles {
			if dashRole == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			response.Error(c, "access denied for your role")
			c.Abort()
			return
		}

		c.Set("userEmail", claims.Email)
		c.Set("userRole", dashRole)

		c.Next()
	}
}
