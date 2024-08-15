package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

type AuthMiddlewareSuite struct {
	suite.Suite
	router     *gin.Engine
	jwtService domain.JwtService
	middleware domain.AuthMiddlerWareInterface
	env        *bootstrap.Env
}

func (suite *AuthMiddlewareSuite) SetupTest() {
	suite.router = gin.Default()
	suite.env = bootstrap.NewEnv()
	suite.jwtService = infrastructure.NewJwtService(suite.env)
	suite.middleware = infrastructure.NewAuthMIddleWare(suite.jwtService)

	// Apply the middleware to a test route
	suite.router.Use(suite.middleware.AuthMiddleware())
	suite.router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	})
}

func (suite *AuthMiddlewareSuite) TestAuthMiddlewareUnauthorized() {
	// Create a request with no Authorization header
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	// Assert that the response status code is 401 Unauthorized
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.Code)
	assert.JSONEq(suite.T(), `{"error": "Unauthorized"}`, resp.Body.String())
}

func (suite *AuthMiddlewareSuite) TestAuthMiddlewareAuthorized() {
	// Create a valid token
	claims := domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: "user@example.com",
		Role:  "user",
	}
	token, _ := suite.jwtService.CreateAccessToken(claims)

	// Create a request with Authorization header
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	// Assert that the response status code is 200 OK
	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.JSONEq(suite.T(), `{"message": "success"}`, resp.Body.String())
}

func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareSuite))
}
