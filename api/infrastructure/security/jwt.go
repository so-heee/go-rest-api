package security

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	jwtExtractor func(echo.Context) (string, error)
)

// Errors
var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
)

// Defaults
const (
	AlgorithmHS256    = "HS256"
	SigningMethod     = AlgorithmHS256
	ContextKey        = "user"
	TokenHeader       = echo.HeaderAuthorization
	AuthScheme        = "Bearer"
	SigningKey        = "secret"
	SigningRefreshKey = "secretRefresh"
)

type JwtClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    int    `json:"user_id"`
	jwt.StandardClaims
}

func VerifyToken() *oapimiddleware.Options {
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
					return nil
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

func VerifyAccessToken(t string) (string, error) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != AlgorithmHS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		return []byte(SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	c := token.Claims.(*JwtClaims)
	return c.Id, nil
}

func VerifyRefreshToken(t string) (string, error) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != AlgorithmHS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		return []byte(SigningRefreshKey), nil
	})
	if err != nil {
		return "", err
	}

	c := token.Claims.(*JwtClaims)
	return c.Id, nil
}

func GenerateAccessToken(id int) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(id, expirationTime, SigningKey)
}

func GenerateRefreshToken(id int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	return generateToken(id, expirationTime, SigningRefreshKey)
}

func generateToken(id int, expirationTime time.Time, key string) (string, error) {

	claims := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(id),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}

// jwtFromQuery returns a `jwtExtractor` that extracts token from the query string.
func jwtFromQuery(param string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		token := c.QueryParam(param)
		if token == "" {
			return "", ErrJWTMissing
		}
		return token, nil
	}
}

// jwtFromParam returns a `jwtExtractor` that extracts token from the url param string.
func jwtFromParam(param string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		token := c.Param(param)
		if token == "" {
			return "", ErrJWTMissing
		}
		return token, nil
	}
}

// jwtFromCookie returns a `jwtExtractor` that extracts token from the named cookie.
func jwtFromCookie(name string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		cookie, err := c.Cookie(name)
		if err != nil {
			return "", ErrJWTMissing
		}
		return cookie.Value, nil
	}
}
