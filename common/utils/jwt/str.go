package jwtUtils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/showurl/Zero-IM-Server/common/types"
	"strings"
	"time"
)

func TokenClaimsString(token string) string {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return ""
	}
	return split[1]
}

type Claims struct {
	UID      string
	Platform string //login platform
	jwt.RegisteredClaims
}

func GetClaimFromToken(tokensString string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, types.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, types.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, types.ErrTokenNotValidYet
			} else {
				return nil, types.ErrTokenUnknown
			}
		} else {
			return nil, types.ErrTokenNotValidYet
		}
	} else {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			//log.NewDebug("", claims.UID, claims.Platform)
			return claims, nil
		}
		return nil, types.ErrTokenNotValidYet
	}
}

func BuildClaims(uid, platform string) Claims {
	now := time.Now()
	return Claims{
		UID:      uid,
		Platform: platform,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 365 * 1000)), //Expiration time
			IssuedAt:  jwt.NewNumericDate(now),                                  //Issuing time
			NotBefore: jwt.NewNumericDate(now),                                  //Begin Effective time
		}}
}
