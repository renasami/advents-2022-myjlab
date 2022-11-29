package auth_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"github.com/renasami/advents-2022-myjlab/api/service/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var identityKey = "email"

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type LoginStruct struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

func HelloHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func Jwt() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Auth); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println(claims)
			return &models.Auth{
				Email:    claims[identityKey].(string),
				Username: "uname",
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginRequest models.LoginRequest
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
			buf := make([]byte, 2048)
			n, _ := c.Request.Body.Read(buf)
			b := buf[0:n]
			//should bind jsonはなんかうまくいかなかった
			if err := json.Unmarshal(b, &loginRequest); err != nil {
				fmt.Println(err)
				c.JSON(400, gin.H{"status": "error", "message": jwt.ErrMissingLoginValues})
				return "", jwt.ErrMissingLoginValues
			}
			fmt.Println(loginRequest)
			email := loginRequest.Email
			result, err := database.GetUserByEmail(email)
			if err != nil {
				c.JSON(400, gin.H{"status": "error", "message": gorm.ErrRecordNotFound})
				return nil, gorm.ErrRecordNotFound
			}
			usr := *result
			pas := []byte(loginRequest.Password)
			err = bcrypt.CompareHashAndPassword([]byte(usr.Password), pas)
			if err != nil {
				log.Error("password is invalid")
				return nil, jwt.ErrFailedAuthentication
			}
			return &models.Auth{
				Email:    email,
				Password: usr.Password,
				Username: usr.Username,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.Auth); ok && v.Email == "drrr0502@gmaili.com" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   token,
				"expire":  t.Format(time.RFC3339),
				"message": "login successfully",
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}
