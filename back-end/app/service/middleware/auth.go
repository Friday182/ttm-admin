package middleware

import (
	"net/http"
	"time"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"

	// "github.com/gogf/gf/util/gvalid"
	"github.com/friday182/ttm-admin/app/model"
	"github.com/jinzhu/gorm"
)

var (
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
	// Customized login parameter validation rules.
	ValidationRules = g.Map{
		"username": "required",
		"password": "required",
	}
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	/*	token := r.Get("token")
		if token == "123456" {
			r.Middleware.Next()
		} else {
			r.Response.WriteStatus(http.StatusForbidden)
		}
	*/
	token := r.GetString("authorization")
	if len(token) > 0 {
		glog.Println("Token is: ", token)
		GfJWTMiddleware.MiddlewareFunc()(r)
		r.Middleware.Next()
	} else {
		glog.Println("Token is Empty!")
		r.Response.WriteJson(g.Map{
			"code": http.StatusForbidden,
			"msg":  "Token is Empty!",
		})
		r.ExitAll()
	}
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	var mentor model.Mentor
	username := r.GetPostString("username")
	passwd := r.GetPostString("password")
	tmpDb := model.GetDb()
	
	err := tmpDb.Where("email = ? AND password = ?", username, passwd).Find(&mentor).Error
	if gorm.IsRecordNotFoundError(err) {
		token = ""
		r.Response.WriteJson(g.Map{
			"code":   http.StatusForbidden,
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	} else {
		r.Response.WriteJson(g.Map{
			"code":   http.StatusOK,
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
		mentor.Token = token
		tmpDb.Save(&mentor)
	}
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {
	username := r.GetPostString("username")
	passwd := r.GetPostString("password")
	if username == "" || passwd == "" {
		return "", jwt.ErrFailedAuthentication
	}
	// if data["username"] == "admin" && data["password"] == "admin" {
	//if model.ValidUser(username, passwd) {
	return g.Map{}, nil
	//}
	//return nil, jwt.ErrFailedAuthentication
}
