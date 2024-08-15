package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

type jwtService struct {
	Env *bootstrap.Env
}

func NewJwtService(env *bootstrap.Env) domain.JwtService {
	return &jwtService{
		Env: env,
	}
}

func (j *jwtService) CreateAccessToken(user domain.UserDTO) (accessToken string, err error) {
	expTime := time.Now().Add(time.Minute * 30).Unix()
	secret := []byte(j.Env.JWT_SECRET)

	claims := &domain.JWTCustome{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return t, err

}


func (j *jwtService) ValidateToken(tokenStr string) (*domain.JWTCustome, error) {
	jwtSecret := []byte(j.Env.JWT_SECRET)

	token, err := jwt.ParseWithClaims(tokenStr, &domain.JWTCustome{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*domain.JWTCustome)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}

func (j *jwtService) ValidateAuthHeader(authHeader string) ([]string, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	return authParts, nil
}

func (j *jwtService) GetClaims(authHeader string) (*domain.JWTCustome, error) {

	// Validate and parse the Authorization header
	authParts, err := j.ValidateAuthHeader(authHeader)
	if err != nil {
		return nil, fmt.Errorf("invalid authorization header: %v", err)
	}

	// Validate the JWT token
	claims, err := j.ValidateToken(authParts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}
