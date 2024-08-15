package infrastructure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	infrastructure "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Infrastructure"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	passwordService domain.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.passwordService = infrastructure.NewPasswordService()
}

func (suite *PasswordServiceTestSuite) TestEncryptPassword() {
	password := "password123"

	encryptedPassword, err := suite.passwordService.EncryptPassword(password)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), encryptedPassword)
}

func (suite *PasswordServiceTestSuite) TestValidatePassword() {
	password := "password123"
	hashedPassword, _ := suite.passwordService.EncryptPassword(password)

	valid := suite.passwordService.ValidatePassword(password, hashedPassword)

	assert.True(suite.T(), valid)
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}
