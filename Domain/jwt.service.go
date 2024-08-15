package domain

type JwtService interface {
	CreateAccessToken(user UserDTO) (accessToken string, err error)
	ValidateToken(tokenStr string) (*JWTCustome, error)
	ValidateAuthHeader(authHeader string) ([]string, error)
	GetClaims(authHeader string) (*JWTCustome, error)
}
