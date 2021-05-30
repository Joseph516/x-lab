package app

import (
	"github.com/dgrijalva/jwt-go"
	"service/global"
	"service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(global.JWTSetting.Expire)

	claims := Claims{
		// AppKey:    util.EncodeMD5(appKey),
		AppKey:    appKey,
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析给定token
func ParseToken(token string) (*Claims, error) {
	// 解析鉴权声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 这样解析会报错
// func GetAppKeyFromToken(token string) (string, error) {
// 	payload, err := base64.StdEncoding.DecodeString(strings.Split(token, ".")[1])
// 	if err != nil {
// 		return "", err
// 	}
// 	claims := Claims{}
// 	err = json.Unmarshal(payload, &claims)
// 	if err != nil {
// 		return "", err
// 	}
// 	return claims.AppKey, nil
// }
