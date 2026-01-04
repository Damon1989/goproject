package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// Package utils contains helper utilities used across the application.
// This file provides JWT token generation utilities.

// JwtCustomClaims defines the custom claims stored in the JWT for this app.
// It embeds jwt.RegisteredClaims to include standard fields like
// ExpiresAt and IssuedAt, and adds application-specific fields ID and Name.
type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// stSigningKey is the HMAC signing key used to sign tokens.
// It is read from configuration (viper) under the key "jwt.signingKey".
// Note: reading the key at package init time means viper must be configured
// before this package is first used, otherwise an empty key may be loaded.
var stSigningKey = []byte(viper.GetString("jwt.signingKey"))

// GenerateToken creates and signs a new JWT for the given user id and name.
// It returns the serialized token string or an error if signing fails.
//
// Behavior details:
//   - The token's Subject is set to "Token".
//   - IssuedAt is set to the current time.
//   - ExpiresAt is set by adding a duration configured at "jwt.tokenExpire".
//     The code multiplies the configured value by time.Minute (matching the
//     original implementation), so ensure the config value represents a
//     numeric duration in minutes. If you store a full duration string
//     (like "15m") in viper, prefer using viper.GetDuration directly.
func GenerateToken(id uint, name string) (string, error) {
	fmt.Println(viper.GetDuration("jwt.tokenExpire"))
	// Build custom claims with standard registered claims
	claims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt is set to now + configured tokenExpire (in minutes).
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * viper.GetDuration("jwt.tokenExpire"))),
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), // for test purpose, set to 30 minutes
			// IssuedAt is the current time
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Subject identifies the purpose of this token
			Subject: "Token",
		},
	}

	// Create a new token object using HMAC SHA-256 and our claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token using the package-level signing key and return the result
	return token.SignedString(stSigningKey)

}

func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	var claims JwtCustomClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
		return stSigningKey, nil
	})
	if err == nil && token != nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return claims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
