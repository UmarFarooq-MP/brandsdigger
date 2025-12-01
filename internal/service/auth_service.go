package service

import "brandsdigger/internal/domain/auth"

type authService struct {
}

func (as *authService) Login(input auth.Login) (string, error) {
	return "", nil
}
func (as *authService) Signup(input auth.SignUp) (string, error) {
	return "", nil
}

func NewAuthService() auth.Auth {
	return &authService{}
}
