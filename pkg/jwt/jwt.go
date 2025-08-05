package jwt

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	SecretKey string
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{SecretKey: secretKey}
}

func (j *Jwt) GenerateToken(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	s, err := t.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return s, nil
}
