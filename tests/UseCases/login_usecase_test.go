package tests

import (
	"context"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
	usecases "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/UseCases"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/constants"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginControllerSuite struct {
	suite.Suite
	repository      *mocks.MockUserRepository
	signupUseCase   domain.SignUpUseCase
	usecase         domain.LoginUseCase
	ctrl            *gomock.Controller
	ctx             context.Context
	ENV             *bootstrap.Env
	passwordService domain.PasswordService
	jwtService      domain.JwtService
}

func (suite *loginControllerSuite) SetupTest() {
	suite.ENV = bootstrap.NewEnv()

	log.Println(suite.ENV.JWT_SECRET)

	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockUserRepository(suite.ctrl)
	suite.passwordService = infrastructure.NewPasswordService()
	suite.jwtService = infrastructure.NewJwtService(suite.ENV)
	suite.signupUseCase = usecases.NewSignUpUseCase(suite.repository, suite.passwordService, suite.jwtService)
	suite.usecase = usecases.NewLoginUseCase(suite.repository, suite.passwordService, suite.jwtService)
}

func (suite *loginControllerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *loginControllerSuite) getUser() (*domain.User, *domain.ErrorResponse) {
	suite.repository.
		EXPECT().
		GetUserEmail(gomock.Any(), constants.TestEmail).
		Return(&domain.User{Email: constants.TestEmail}, nil)

	return suite.usecase.GetUserEmail(suite.ctx, constants.TestEmail)
}

func (suite *loginControllerSuite) validatePassword(password, hashedPassword string, expected bool) {
	valid := suite.usecase.ValidatePassword(password, hashedPassword)
	suite.Equal(expected, valid, "Password validation result mismatch")
}

func (suite *loginControllerSuite) hashPassword(password string) (string, error) {
	return suite.signupUseCase.EncryptPassword(password)
}

func (suite *loginControllerSuite) createTestUser(err *domain.ErrorResponse) (domain.UserDTO, *domain.ErrorResponse) {

	userReq := domain.UserCreateRequest{
		Email:    constants.TestEmail,
		Password: constants.TestPassword,
	}

	createdUser := domain.UserDTO{
		ID:    primitive.NewObjectID(),
		Email: userReq.Email,
		Role:  "admin",
	}

	suite.repository.EXPECT().
		CreateUser(gomock.Any(), userReq).
		Return(createdUser, err).
		Times(1)

	return suite.signupUseCase.CreateUser(suite.ctx, userReq)
}

func (suite *loginControllerSuite) TestGetUserEmail() {
	_, err := suite.createTestUser(nil)
	suite.Nil(err, "Error creating user")

	user, retErr := suite.getUser()
	suite.Nil(retErr, "Error retrieving user")
	suite.NotEmpty(user, "The retrieved user shouldn't be empty")

}

func (suite *loginControllerSuite) TestGetUserEmail_NotFound() {
	suite.repository.
		EXPECT().
		GetUserEmail(gomock.Any(), constants.InvalidEmail).
		Return(&domain.User{}, &domain.ErrorResponse{Message: "User not found"})

	retrievedUser, retErr := suite.usecase.GetUserEmail(suite.ctx, constants.InvalidEmail)

	suite.Error(retErr, "Expected error not received")
	suite.Equal(&domain.User{}, retrievedUser, "Retrieved user should be empty")
	suite.Contains(retErr.Message, "User not found", "Error message mismatch")
}

func (suite *loginControllerSuite) TestValidatePassword_Valid() {
	hashedPassword, err := suite.hashPassword(constants.TestPassword)
	suite.Nil(err, "Error hashing password")

	suite.validatePassword(constants.TestPassword, hashedPassword, true)
}

func (suite *loginControllerSuite) TestValidatePassword_Invalid() {
	hashedPassword, err := suite.hashPassword(constants.TestPassword)
	suite.Nil(err, "Error hashing password")

	suite.validatePassword(constants.InvalidPassword, hashedPassword, false)
}

func TestLoginUseCase(t *testing.T) {
	suite.Run(t, new(loginControllerSuite))
}
