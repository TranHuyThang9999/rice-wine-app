package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/services"
	"strings"
)

type Middleware struct {
	jwt         *services.JWTService
	userService *services.UserService
}

func NewMiddleware(jwt *services.JWTService,
	userService *services.UserService) *Middleware {
	return &Middleware{jwt: jwt, userService: userService}
}
func (u *Middleware) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		userClaims, err := u.jwt.Verify(c, tokenString)
		if err != nil {
			switch {
			case errors.Is(err, apperrors.ErrInvalidToken), errors.Is(err, apperrors.ErrTokenExpired), errors.Is(err, apperrors.ErrUnexpectedSigningMethod):
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to authenticate token"})
			}
			c.Abort()
			return
		}
		info, err := u.userService.ProFileUser(c, userClaims.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to authenticate token"})
			c.Abort()
			return
		}
		if info.UpdatedAt != userClaims.UpdateAt {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to authenticate token"})
			c.Abort()
		}
		c.Set("userID", info.ID)
		c.Set("role", userClaims.Role)
		c.Set("phoneNumber", userClaims.PhoneNumber)
		c.Next()
	}
}
