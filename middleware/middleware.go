package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CheckPermissions(path string, requiredPermissions []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userPermissions, err := getUserPermissions(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if !hasPermission(userPermissions, policy{
				Path:         path,
				Capabilities: requiredPermissions,
			}) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getUserPermissions(r *http.Request) ([]policy, error) {
	tokenString := r.Header.Get("X-Vault-Token")
	if tokenString == "" {
		return nil, fmt.Errorf("missing token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("security.signing_key")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Policies, nil
	}

	return nil, fmt.Errorf("invalid token")
}

type policy struct {
	Path         string   `json:"path"`
	Capabilities []string `json:"capabilities"`
}
type tokenClaims struct {
	jwt.StandardClaims
	Policies []policy `json:"policies"`
}

func hasPermission(userPolicies []policy, requiredPolicy policy) bool {
	for i := range userPolicies {
		if userPolicies[i].Path == requiredPolicy.Path &&
			slicesEqual(userPolicies[i].Capabilities, requiredPolicy.Capabilities) {
			return true
		}
	}
	return false
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[string]int)

	for _, item := range a {
		count[item]++
	}

	for _, item := range b {
		count[item]--
		if count[item] < 0 {
			return false
		}
	}

	return true
}
