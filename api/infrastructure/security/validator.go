package security

import (
	"context"
	"fmt"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidatiorOptions() *oapimiddleware.Options {
	return &oapimiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {

				if input.SecuritySchemeName != "BearerAuth" {
					return &echo.HTTPError{
						Code:    ErrJWTInvalid.Code,
						Message: ErrJWTInvalid.Message,
					}
				}

				ec := oapimiddleware.GetEchoContext(c)
				if ec == nil {
					return echo.ErrBadRequest
				}

				extractor := jwtFromHeader(echo.HeaderAuthorization, AuthScheme)
				auth, err := extractor(ec)
				if err != nil {
					return err
				}

				claims := &JwtClaims{}

				token, err := jwt.ParseWithClaims(auth, claims, func(t *jwt.Token) (interface{}, error) {
					// Check the signing method
					if t.Method.Alg() != AlgorithmHS256 {
						return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
					}

					return []byte("secret"), nil
				})
				if err == nil && token.Valid {
					// claims.UserID = u.ID
					// claims.UserID = u
					// Store user information from token into context.
					ec.Set(ContextKey, claims)
				}

				return &echo.HTTPError{
					Code:     ErrJWTInvalid.Code,
					Message:  ErrJWTInvalid.Message,
					Internal: err,
				}
			},
		},
	}
}
