package router

import (
	"github.com/friday182/ttm-admin/app/api/gql"
	ws "github.com/friday182/ttm-admin/app/api/websocket"
	"github.com/friday182/ttm-admin/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const VERSION = "V0.1.0"

func init() {
	s := g.Server()

	/* EnablePProf will define the following routes for debug
	  /debug/pprof/*action
		/debug/pprof/cmdline
		/debug/pprof/profile
		/debug/pprof/symbol
		/debug/pprof/trace
	*/
	s.EnablePProf()
	s.EnableAdmin("/admin")

	// Config server
	// s.SetDumpRouteMap(false)

	// Some IE always request favicon.ico
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// s.BindHandler("/ws", ws.Handler)
	//s.BindHandler("/query", gql.GraphqlHandler)
	//s.BindHandler("/graphql", gql.GraphqlHandler)
	//s.BindHandler("/playground", gql.PlaygroundHandler)

	// Register routes by group
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS)
		g.ALL("/version", func(r *ghttp.Request) { r.Response.WriteExit(VERSION) })
		g.ALL("/query", gql.GraphqlHandler)
		g.ALL("/graphql", gql.GraphqlHandler)
		g.ALL("/playground", gql.PlaygroundHandler)
	})

	s.Group("/api", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS) //, middleware.Auth)
		g.ALL("/refresh_token", middleware.GfJWTMiddleware.RefreshHandler)
		g.ALL("/ws", ws.Handler)
	})

	s.SetPort(8085)
}
