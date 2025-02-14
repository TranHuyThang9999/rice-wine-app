package middlewares

import (
	"errors"
	"net/http"
	"rice-wine-shop/common/log"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/enums"
	"rice-wine-shop/core/services"
	"strings"

	"github.com/gin-gonic/gin"
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
			log.Info("Authorization header is required")
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
		c.Set("role", info.Role)
		c.Set("phoneNumber", userClaims.PhoneNumber)
		c.Next()
	}
}

func (u *Middleware) CheckToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		log.Info("Authorization header is required")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		ctx.Abort()
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		ctx.Abort()
		return
	}

	tokenString := tokenParts[1]
	userClaims, err := u.jwt.Verify(ctx, tokenString)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidToken), errors.Is(err, apperrors.ErrTokenExpired), errors.Is(err, apperrors.ErrUnexpectedSigningMethod):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate token"})
		}
		ctx.Abort()
		return
	}
	info, err := u.userService.ProFileUser(ctx, userClaims.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to authenticate token"})
		ctx.Abort()
		return
	}
	if info.UpdatedAt != userClaims.UpdateAt {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to authenticate token"})
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"userId":       info.ID,
			"role":         info.Role,
			"phone_number": info.PhoneNumber,
		},
	})
}

func (u *Middleware) AuthorizationAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Info("Authorization header is required")
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
			return
		}
		if info.Role != enums.ROLE_ADMIN {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}
		c.Set("userID", info.ID)
		c.Set("role", info.Role)
		c.Set("phoneNumber", userClaims.PhoneNumber)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}
