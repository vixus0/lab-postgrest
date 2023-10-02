package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PostgrestClaims struct {
	Role string `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	secret := flag.String("secret", "./secret.txt", "signing secret file")
	role := flag.String("role", "api_anon", "jwt role")
	username := flag.String("username", "anonymous", "jwt username")

	flag.Parse()
	
	claims := PostgrestClaims{
		Role: *role,
		Username: *username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer: "test",
			Audience: jwt.ClaimStrings{"postgrest"},
		},
	}

	secretBytes, err := os.ReadFile(*secret)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(ss)
}
