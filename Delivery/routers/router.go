package routers

import (
	"github.com/gin-gonic/gin"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine) {

	publicRouter := gin.Group("")
	NewSignupRouter(env, db, publicRouter)
	NewLoginRouter(env, db, publicRouter)
	jwtService := infrastructure.NewJwtService(env)
	authMiddleware := infrastructure.NewAuthMIddleWare(jwtService)

	protectedRouter := gin.Group("")
	protectedRouter.Use(authMiddleware.AuthMiddleware())

	adminGroup := protectedRouter.Group("")
	adminGroup.Use(authMiddleware.RoleBasedMiddleWare(env.ALLOWED_USERS))

	NewPromoteRouter(env, db, adminGroup)
	NewTaskRouter(env, db, adminGroup)
}
