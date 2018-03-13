package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	// TokenExpired Token过期
	TokenExpired = errors.New("Token is expired")
	// TokenNotValidYet Token尚未生效
	TokenNotValidYet = errors.New("Token not active yet")
	// TokenMalformed Token异常
	TokenMalformed = errors.New("That's not even a token")
	// TokenInvalid Token无效
	TokenInvalid = errors.New("Couldn't handle this token")
	// SignKey 签署密钥
	SignKey = "secret"
)

// JWT JSON Web Token
type JWT struct {
	SigningKey []byte
}

// CustomClaims .
type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// NewJWT .
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// SetSignKey .
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// GetSignKey .
func GetSignKey() string {
	return SignKey
}

// CreateToken 创建Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			}
		}
		return nil, TokenInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// JWTAuth 处理认证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("Authorization", "")
		if token == "" {
			if token, err := c.Cookie("Authorization"); err == nil {
				if s := strings.Split(token, " "); len(s) == 2 {
					token = s[1]
				}
			}
		}

		j := NewJWT()

		claims, err := j.ParseToken(token)

		if err != nil {
			if err == TokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.Header("Authorization", "Bear "+token)
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": 1, "message": err.Error()})
			return
		}
		c.Set("claims", claims)
	}
}
