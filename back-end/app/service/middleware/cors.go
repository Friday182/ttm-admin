package middleware

import "github.com/gogf/gf/net/ghttp"

// 允许接口跨域请求
func CORS(r *ghttp.Request) {
	//corsOptions := r.Response.DefaultCORSOptions()
	//corsOptions.AllowDomain = []string{"*", "localhost"}
	//r.Response.CORS(corsOptions)
	r.Response.CORSDefault()
	r.Middleware.Next()
}