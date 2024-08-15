package infrastructure

import (
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() domain.PasswordService {
	return &passwordService{}
}

func (p *passwordService) EncryptPassword(password string) (string, error) {
	cur_pass := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(cur_pass, bcrypt.DefaultCost)

	return string(encryptedPassword), err

}

func (p *passwordService) ValidatePassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
