package jwt

import "github.com/golang-jwt/jwt/v5"

type JwtData struct {
	Email string
}

type Jwt struct {
	SecretKey string
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{SecretKey: secretKey}
}

func (j *Jwt) GenerateToken(data JwtData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	s, err := t.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *Jwt) Parse(token string) (bool, *JwtData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JwtData{
		Email: email.(string),
	}
}
