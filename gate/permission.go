package gate

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthInfo struct {
	ID    string   `json:"id"`
	Role  string   `json:"role"`
	Scope []string `json:"scope"`
}

func (a *AuthInfo) HasRole(role ...string) bool {
	for _, vRole := range role {
		if strings.EqualFold(vRole, a.Role) {
			return true
		}
	}

	return false
}

func (a *AuthInfo) HasScope(permission ...string) bool {
	for _, vScope := range a.Scope {
		for _, vPer := range permission {
			if strings.EqualFold(vPer, vScope) {
				return true
			}
		}
	}
	return false
}

func GetAuthInfo(secretKey, stringToken string, fn func(claims jwt.MapClaims) (AuthInfo, error)) (AuthInfo, error) {

	var result AuthInfo

	if secretKey == "" {
		return result, fmt.Errorf(`not found secretKey`)
	}

	// ======= 2. แปลงกลับ (Parse) =======
	parsedToken, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", t.Header["alg"])

		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return result, err
	}

	if !parsedToken.Valid {
		return result, fmt.Errorf("invalid or expired token")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {

		resp, err := fn(claims)
		if err != nil {
			return result, err
		}

		result = resp

	} else {
		return result, fmt.Errorf("invalid token")
	}

	return result, nil
}
