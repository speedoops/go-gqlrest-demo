package errorsx

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// InvalidParamError 表达无效参数类错误
type InvalidParamError struct {
	err error
}

func (e InvalidParamError) Error() string {
	return "INLALID_PARAM: " + e.err.Error()
}

// NewInvalidParamError return a InvalidParamError instance
func NewInvalidParamError(err error) error {
	return &InvalidParamError{
		err: err,
	}
}

// ResolverError 表达GraphL解析器错误
type ResolverError struct {
	err error
}

func (e ResolverError) Error() string {
	return "INTERNAL_ERROR: " + e.err.Error()
}

// NewResolverError return a ResolverError instance
func NewResolverError(err error) error {
	return &ResolverError{
		err: err,
	}
}

// MyErrorPresenter 将自定义错误友好呈现给客户端
func MyErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := new(gqlerror.Error)

	var paramErr InvalidParamError
	if errors.As(e, &paramErr) {
		err.Message = "InvalidParamError: " + paramErr.Error()
		return err
	}

	// if pe, ok := e.(*InvalidParamError); ok {
	// 	err.Message = "InvalidParamError: " + pe.Error()
	// 	return err
	// }

	var resolverErr ResolverError
	if errors.As(e, &resolverErr) {
		err.Message = "GraphResolverError: " + resolverErr.Error()
		return err
	}

	// var httpParamErr httpclient.InvalidParamError
	// if errors.As(e, &httpParamErr) {
	// 	err.Message = "HttpInvalidParamError: " + httpParamErr.Error()
	// 	return err
	// }

	// var httpAuthFailureErr httpclient.AuthFailureError
	// if errors.As(e, &httpAuthFailureErr) {
	// 	err.Message = "HttpAuthFailureError: " + httpAuthFailureErr.Error()
	// 	return err
	// }

	// var httpRequestErr httpclient.SwaggerRequestError
	// if errors.As(e, &httpRequestErr) {
	// 	err.Message = "HttpRequestError: " + httpRequestErr.Error()
	// 	return err
	// }

	// var httpResponseErr httpclient.SwaggerResponseError
	// if errors.As(e, &httpResponseErr) {
	// 	err.Message = "HttpResponseError: " + httpResponseErr.Error()
	// 	return err
	// }

	// var httpDataErr httpclient.SwaggerDataError
	// if errors.As(e, &httpDataErr) {
	// 	err.Message = "HttpDataError: " + httpDataErr.Error()
	// 	return err
	// }

	errcode.Set(err, "500")
	err.Message = fmt.Sprintf("UnknownError: %s !!!FIXME!!! %#v", err.Message, e)
	return err
}
