package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paularah/wallet/pkg/jwt"
)

const (
	authHeaderKey = "authorization"
	claimKey      = "user"
)

func authMiddleware(tokener jwt.Tokener) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		jwtToken, err := validateBearerAuthHeader(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		claim, err := tokener.VerifyJWT(jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(claimKey, claim)
		ctx.Next()

	}
}

func validateBearerAuthHeader(authHeader string) (string, error) {
	var err error
	if len(authHeader) == 0 {
		err = errors.New("authorization header is missing")
		return "", err
	}

	fields := strings.Fields(authHeader)

	if len(fields) < 2 {
		err = errors.New("invalid authorization header format")
		return "", err
	}

	return fields[1], err
}
