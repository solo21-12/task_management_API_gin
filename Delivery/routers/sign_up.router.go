package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Delivery/controllers"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	repositories "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Repositories"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignupRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {

	userRepo := repositories.NewUserRepository(db, env.USER_COLLECTION)
	passwordService := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJwtService(env)
	userCase := usecases.NewSignUpUseCase(userRepo, passwordService, jwtService)
	signUpController := controllers.SignupController{
		SignupUsecase: userCase,
		Env:           env,
	}

	group.POST("/register", signUpController.SignUp)

}
