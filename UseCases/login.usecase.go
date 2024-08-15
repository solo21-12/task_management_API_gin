package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

type loginUseCase struct {
	userRespository domain.UserRepository
	passwordService domain.PasswordService
	jwtService      domain.JwtService
}

func NewLoginUseCase(userRespository domain.UserRepository, passwordService domain.PasswordService, jwtService domain.JwtService) domain.LoginUseCase {
	return &loginUseCase{
		userRespository: userRespository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (l *loginUseCase) GetUserEmail(ctx context.Context, email string) (*domain.User, *domain.ErrorResponse) {
	return l.userRespository.GetUserEmail(ctx, email)

}
func (l *loginUseCase) CreateAccessToken(user domain.UserDTO) (accessToken string, err error) {
	return l.jwtService.CreateAccessToken(user)

}
func (l *loginUseCase) ValidatePassword(password string, hashedPassword string) bool {
	return l.passwordService.ValidatePassword(password, hashedPassword)
}
