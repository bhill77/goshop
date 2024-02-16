package middleware

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func (JwtCustomClaims) Valid() error {
	return nil
}

// jwtConfig := echojwt.Config{
// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 		return new(JwtCustomClaims)
// 	},
// 	SigningKey: []byte(conf.JwtSecret),
// }
