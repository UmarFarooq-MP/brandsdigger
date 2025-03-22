package auth

type TokenService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}
