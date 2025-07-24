package Infrastructure

import "golang.org/x/crypto/bcrypt"

// PasswordService defines methods for password hashing
type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashed, password string) error
}

// passwordService implemnts PasswordService
type passwordService struct{}

// ComparePassword implements PasswordService.
func (p *passwordService) ComparePassword(hashed string, password string) error {
	return  bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

// HashPassword implements PasswordService.
func (p *passwordService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hashed), nil
}

// NewPasswordService creates a new PasswordService
func NewPasswordService() PasswordService {
	return &passwordService{}
}
