package infrastructure_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

type JwtServiceSuite struct {
	suite.Suite
	jwtSvc domain.JwtService
}

func (suite *JwtServiceSuite) SetupTest() {
	env := &bootstrap.Env{
		JWT_SECRET: "secret",
	}
	suite.jwtSvc = infrastructure.NewJwtService(env)
}

func (suite *JwtServiceSuite) TestCreateAccessToken() {
	user := createTestUser()

	accessToken, err := suite.jwtSvc.CreateAccessToken(user)

	suite.assertNoError(err)
	suite.NotEmpty(accessToken)
}

func (suite *JwtServiceSuite) TestValidateToken() {
	accessToken := suite.createTestAccessToken()

	claims, err := suite.jwtSvc.ValidateToken(accessToken)

	suite.assertNoError(err)
	suite.NotNil(claims)
}

func (suite *JwtServiceSuite) TestValidateAuthHeader() {
	authHeader := "Bearer valid_token"

	authParts, err := suite.jwtSvc.ValidateAuthHeader(authHeader)

	suite.assertNoError(err)
	suite.NotNil(authParts)
}

func (suite *JwtServiceSuite) TestGetClaims() {
	accessToken := suite.createTestAccessToken()
	authHeader := "Bearer " + accessToken

	claims, err := suite.jwtSvc.GetClaims(authHeader)

	suite.assertNoError(err)
	suite.NotNil(claims)
}

func (suite *JwtServiceSuite) assertNoError(err error) {
	if err != nil {
		suite.FailNow("Error: " + err.Error())
	}
}

func createTestUser() domain.UserDTO {
	return domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Role:  "admin",
	}
}

func (suite *JwtServiceSuite) createTestAccessToken() string {
	user := createTestUser()

	accessToken, _ := suite.jwtSvc.CreateAccessToken(user)

	return accessToken
}

func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}
