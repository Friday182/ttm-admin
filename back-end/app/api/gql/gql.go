package gql

import (
	"github.com/99designs/gqlgen/handler"
	//"github.com/jinzhu/gorm"
	//"net/http"
	"github.com/friday182/ttm-admin/app/model"
	gen "github.com/friday182/ttm-admin/app/service/graphql/generated"
	res "github.com/friday182/ttm-admin/app/service/graphql/resolvers"

	//"github.com/friday182/ttm-admin/app/service/middleware"
	"github.com/gogf/gf/net/ghttp"
)

func PlaygroundHandler(r *ghttp.Request) {
	h := handler.Playground("GraphQL", "/query")

	//h.ServeHTTP(r.Response.Writer.RawWriter(), r.Request)
	//return func(*http.HandlerFunc) {
	//	h.ServeHTTP(r.Response.ResponseWriter.RawWriter(), r.Request)
	//}
	h.ServeHTTP(r.Response.ResponseWriter.RawWriter(), r.Request)
}

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(r *ghttp.Request) {

	// Verify token
	//middleware.Auth(r)

	cfg := gen.Config{
		Resolvers: &res.Resolver{
			TtmDb:  model.GetTtmDb(),
			QuizDb: model.GetQuizDb(),
		},
	}
	h := handler.GraphQL(gen.NewExecutableSchema(cfg))
	// )

	h.ServeHTTP(r.Response.ResponseWriter.RawWriter(), r.Request)

	//return func(r *ghttp.Request) {
	//	h.ServeHTTP(r.Response.Writer, r)
	//}
}
