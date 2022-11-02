package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Claim struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func (claim *Claim) Valid() error {
	if time.Now().After(claim.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}

func CreateClaim(userID int64, duration time.Duration) (*Claim, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	claim := &Claim{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return claim, err
}

func CreateJWTToken(userID int64, duration time.Duration, secretKey string) (string, *Claim, error) {
	claim, err := CreateClaim(userID, duration)
	if err != nil {
		return "", nil, nil
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, err
	}
	return token, claim, nil

}

func VerifyJWT(token string, secretKey string) (*Claim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claim{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claim, ok := jwtToken.Claims.(*Claim)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claim, nil

}
