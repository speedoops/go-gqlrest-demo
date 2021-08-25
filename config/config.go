package config

import (
	"github.com/tal-tech/go-zero/rest"
)

// Config 是本应用的配置类
type Config struct {
	rest.RestConf
	GraphQL GraphQLConf
	//Backend httpclient.Backend
}

type GraphQLConf struct {
	Debug struct {
		EnableVerbose bool `json:",optional"` //nolint:revive,staticcheck // FIXME: go-zero
	}
}

// TODO: 破坏了代码分层，后续优化
var GraphQL GraphQLConf
