package mock

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/speedoops/go-gqlrest-demo/config"
	"github.com/speedoops/go-gqlrest-demo/graph/errorsx"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/speedoops/go-gqlrest-demo/graph/model"
	"github.com/speedoops/go-gqlrest/handlerx"
	"github.com/tal-tech/go-zero/core/logx"
)

// FindConfigFile 查找并返回 config.yaml 的完整路径
func FindConfigFile() string {
	_, filename, _, _ := runtime.Caller(0)
	lastDir := path.Dir(filename)
	fileName := "config.yaml"
	for {
		currentPath := fmt.Sprintf("%s/%s", lastDir, fileName)
		if fi, err := os.Stat(currentPath); err == nil {
			if mode := fi.Mode(); mode.IsRegular() {
				return currentPath
			}
		}

		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			panic("not found")
		}
		lastDir = newDir
	}
}

func NewGraphQLServer(resolver generated.ResolverRoot) http.Handler {
	// 1. 初始化服务端配置
	var c config.Config
	//conf.MustLoad(FindConfigFile(), &c)
	config.GraphQL = c.GraphQL
	c.Log.Mode = "console"
	logx.MustSetup(c.Log)

	// 2. 运行 GraphQL Server
	cfg := generated.Config{Resolvers: resolver}
	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		role model.Role) (interface{}, error) {
		// if !getCurrentUser(ctx).HasRole(role) {
		// 	// block calling the next resolver
		// 	return nil, fmt.Errorf("Access denied")
		// }
		log.Println("hasRole")

		// or let it pass through
		return next(ctx)
	}
	cfg.Directives.Hide = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		forArg []string) (res interface{}, err error) {
		return next(ctx)
	}
	cfg.Directives.Http = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		url string, method *string) (res interface{}, err error) {
		return next(ctx)
	}
	cfg.Directives.Preview = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		toggledBy string) (res interface{}, err error) {
		return next(ctx)
	}

	mux := chi.NewRouter()
	srv := handlerx.NewDefaultServer(generated.NewExecutableSchema(cfg))
	srv.SetErrorPresenter(errorsx.MyErrorPresenter)
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logx.ErrorStack("internal server error")
		return errors.New("internal server error")
	})
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	generated.RegisterHandlers("", mux, srv)

	return &Server{mux: mux, srv: srv}
}

type (
	Server struct {
		mux *chi.Mux
		srv *handler.Server
	}
)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
