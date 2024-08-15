package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

const (
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

type authMiddleWare struct {
	jwtService domain.JwtService
}

func NewAuthMIddleWare(jwtService domain.JwtService) domain.AuthMiddlerWareInterface {
	return &authMiddleWare{
		jwtService: jwtService,
	}
}

func (a *authMiddleWare) AuthMiddleware() gin.HandlerFunc {
	// this middleware checks if the user is authenticated
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		_, err := a.jwtService.GetClaims(authHeader)

		if err != nil {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (a *authMiddleWare) RoleBasedMiddleWare(roles ...string) gin.HandlerFunc {
	// this middleware checks if the user is authorized to perform an action
	return func(ctx *gin.Context) {

		claims, err := a.jwtService.GetClaims(ctx.GetHeader("Authorization"))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		userRole := claims.Role

		method := ctx.Request.Method

		switch method {
		case POST, DELETE, PUT, PATCH:
			authrized := false

			for _, role := range roles {
				if role == userRole {
					authrized = true
					break
				}
			}

			if !authrized {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden only admin user can perform this action"})
				ctx.Abort()
				return
			}
		}

		ctx.Next()

	}
}
