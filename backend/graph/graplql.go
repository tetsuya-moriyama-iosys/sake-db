package graph

import (
	"backend/graph/generated"
	"backend/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphQLServer(resolver *resolver.Resolver) *handler.Server {
	//srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	// GraphQLサーバーをセットアップ
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
		Directives: generated.DirectiveRoot{
			Auth: authDirective, // @authディレクティブを適用
		},
	}))

	// POSTトランスポートを追加
	srv.AddTransport(transport.POST{})

	// 必要に応じてGETとOPTIONSもサポート
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Options{})

	// Introspectionを有効にする（GraphiQLからのクエリのため）
	srv.Use(extension.Introspection{})

	return srv
}
