module github.com/vektah/gqlgen-todos

go 1.14

require (
	github.com/99designs/gqlgen v0.13.1-0.20210729011107-9a214e80158b
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/go-chi/chi/v5 v5.0.3
	github.com/go-resty/resty/v2 v2.6.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/speedoops/gql2rest v0.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tal-tech/go-zero v1.1.10
	github.com/vektah/gqlparser/v2 v2.2.0
	golang.org/x/sys v0.0.0-20210816032535-30e4713e60e3 // indirect
)

// replace github.com/99designs/gqlgen => ../../gqlgen
// replace github.com/speedoops/gql2rest => ../../gql2rest
