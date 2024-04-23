package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new acess token for a specific
	CreateToken(username string) (string, *Payload, error)

	// CreateRefresh creates a new refresh token for a specific
	CreateRefreshToken(username string) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey       string
	duration        time.Duration
	refreshDuration time.Duration
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string, duration time.Duration, refreshDuration time.Duration) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey, duration, refreshDuration}, nil
}

// CreateToken creates a new token for a specific
func (maker *JWTMaker) CreateToken(username string) (string, *Payload, error) {
	payload, err := NewPayload(username, maker.duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) CreateRefreshToken(username string) (string, *Payload, error) {
	payload, err := NewPayload(username, maker.refreshDuration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
