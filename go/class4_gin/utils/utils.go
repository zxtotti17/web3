package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"class4_gin/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Valid implements jwt.Claims.
// func (c *CustomClaims) Valid() error {
// 	panic("unimplemented")
// }

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func LoginToken(user models.User, c *gin.Context) (claims models.CustomClaims, t string) {
	claims = models.CustomClaims{
		Uid:      user.ID,
		UserName: user.UserName,
		Exp:      time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
			IssuedAt:  time.Now().Unix(),
			Issuer:    "class4_gin",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(models.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	SetToken(c, tokenString, int(time.Hour*24))
	return claims, tokenString
}

func InitLogger() (l *zap.Logger) {
	l, _ = zap.NewDevelopment()

	defer l.Sync()
	return l
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": message})
}

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := GetToken(c)
		if token == "" {
			NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		fmt.Println("中间件执行", err, claims)
		if err != nil || claims == nil {
			NoAuth(err.Error(), c)
			ClearToken(c)
			c.Abort()
			return
		} else {
			if claims.Exp < time.Now().Unix() {
				NoAuth("授权已过期", c)
				ClearToken(c)
				c.Abort()
				return
			}
			if claims.Uid == 0 {
				NoAuth("用户不存在", c)
				ClearToken(c)
				c.Abort()
				return
			}
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		models.Logger.Info("claims", zap.Any("claims", claims))
		c.Set("claims", claims)
		c.Next()

		// if newToken, exists := c.Get("new-token"); exists {
		// 	c.Header("new-token", newToken.(string))
		// }
		// if newExpiresAt, exists := c.Get("new-expires-at"); exists {
		// 	c.Header("new-expires-at", newExpiresAt.(string))
		// }
	}
}
