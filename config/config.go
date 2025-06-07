package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf // go-zero 的依赖是不必要的，这里只是为了说明如何集成
	GraphQL       GraphQLConf
}

type GraphQLConf struct {
	Debug struct {
		EnableVerbose bool `json:",optional"` //nolint:revive,staticcheck
	}
}

// TODO: 破坏了代码分层，后续优化
var GraphQL GraphQLConf
