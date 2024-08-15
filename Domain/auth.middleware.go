package domain

import "github.com/gin-gonic/gin"

type AuthMiddlerWareInterface interface {
	AuthMiddleware() gin.HandlerFunc
	RoleBasedMiddleWare(roles ...string) gin.HandlerFunc
}
