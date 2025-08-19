package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_super_secret_key_that_is_very_long_and_secure")

type JWTClaim struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, roleID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization is not valid"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "InvalidFormat"})
			return
		}
		claims := &JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("beklenmedik imzalama metodu: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Jeton"})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("role_id", claims.RoleID)
		ctx.Next()
	}
}
func AuthAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID_interface, exists := c.Get("role_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Role Problem"})
			return
		}
		roleID, ok := roleID_interface.(uint)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Role Format"})
			return
		}

		const adminRoleID uint = 0

		if roleID != adminRoleID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Bu işlemi yapmak için yönetici yetkisine sahip olmalısınız"})
			return
		}
		c.Next()
	}
}
