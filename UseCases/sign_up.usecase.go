package usecases

import (
	"context"

	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

type signUpUseCase struct {
	userRespository domain.UserRepository
	passwordService domain.PasswordService
	jwtService      domain.JwtService
}

func NewSignUpUseCase(userRespository domain.UserRepository, passwordService domain.PasswordService, jwtService domain.JwtService) domain.SignUpUseCase {
	return &signUpUseCase{
		userRespository: userRespository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (s *signUpUseCase) CreateUser(ctx context.Context, user domain.UserCreateRequest) (domain.UserDTO, *domain.ErrorResponse) {
	return s.userRespository.CreateUser(ctx, user)
}

func (s *signUpUseCase) CreateAccessToken(user domain.UserDTO) (accessToken string, err error) {
	return s.jwtService.CreateAccessToken(user)
}

func (s *signUpUseCase) EncryptPassword(password string) (string, error) {
	return s.passwordService.EncryptPassword(password)
}
