package monimeapis

import (
	"context"
	"errors"
	"fmt"
	"github.com/monime-lab/gok/syserr"
	"github.com/monime-lab/gwater/httputil"
	"github.com/monime-lab/gwater/mimetype"
	webclient "github.com/monime-lab/web-client-go"
	"github.com/monimesl/monime-cli/cli-utils/monimeapis"
	"net/http"
	"reflect"
)

type requestOptionConfig[Resp any] struct {
	skipTokenPlacement bool
	rwb                webclient.RequestWithBody
	apiErrorMapper     func(apiErr ApiError) (error, bool)
}

type RequestOption[Resp any] interface {
	apply(c *requestOptionConfig[Resp]) webclient.RequestWithBody
}

func SkipTokenPlacement[Resp any](skip bool) RequestOption[Resp] {
	return requestFunc[Resp](func(c *requestOptionConfig[Resp]) webclient.RequestWithBody {
		c.skipTokenPlacement = skip
		return c.rwb
	})
}

func WithApiErrorMapper[Resp any](mapper func(apiError ApiError) (error, bool)) RequestOption[Resp] {
	return requestFunc[Resp](func(c *requestOptionConfig[Resp]) webclient.RequestWithBody {
		c.apiErrorMapper = mapper
		return c.rwb
	})
}

func WithRequestOptionFunc[Resp any](f func(r webclient.RequestWithBody) webclient.RequestWithBody) RequestOption[Resp] {
	return requestFunc[Resp](func(c *requestOptionConfig[Resp]) webclient.RequestWithBody {
		return f(c.rwb)
	})
}

type requestFunc[Resp any] func(c *requestOptionConfig[Resp]) webclient.RequestWithBody

//nolint:unused
func (f requestFunc[Resp]) apply(c *requestOptionConfig[Resp]) webclient.RequestWithBody {
	return f(c)
}

type ApiResult[T any] struct {
	Success  bool          `json:"success"`
	Messages []interface{} `json:"messages"`
	Result   T             `json:"result"`
	// states
	StatusCode int         `json:"-"`
	Headers    http.Header `json:"-"`
}

type ApiError struct {
	Success  bool          `json:"success"`
	Messages []interface{} `json:"messages"`
	Error_   struct {
		Code    int           `json:"code"`
		Reason  string        `json:"reason"`
		Message string        `json:"message"`
		Details []interface{} `json:"details"`
	} `json:"error"`
	StatusCode int         `json:"-"`
	Headers    http.Header `json:"-"`
}

func (e ApiError) Error() string {
	return e.Error_.Message
}

func ApiRequest[Req, T any](ctx context.Context, client webclient.Client, method, url string, body Req, options ...RequestOption[T]) (*ApiResult[T], error) {
	config, result, err := apiRequest[Req, T](ctx, client, method, url, body, options...)
	if err == nil {
		return result, nil
	}
	var apiErr ApiError
	ok := errors.As(err, &apiErr)
	if !ok {
		return result, err
	}
	if config.apiErrorMapper != nil {
		if newErr, ok := config.apiErrorMapper(apiErr); ok {
			return result, newErr
		}
	}
	switch apiErr.StatusCode {
	case http.StatusUnauthorized:
		return result, monimeapis.ErrNotAuthenticated
	case http.StatusNotFound:
		return result, syserr.ResourceMissing(apiErr.Error())
	case http.StatusConflict:
		return result, syserr.ResourceConflicted(apiErr.Error())
	case http.StatusBadRequest:
		return result, syserr.InvalidRequest(apiErr.Error())
	}
	return result, err
}

//nolint:cyclop
func apiRequest[Req any, T any](ctx context.Context, client webclient.Client, method, url string, body Req, options ...RequestOption[T]) (*requestOptionConfig[T], *ApiResult[T], error) {
	if client == nil {
		client = getWebClient()
	}
	rwb := client.RequestWithBody(method, url).SetContext(ctx)
	if !isNil(body) {
		rwb = rwb.SetBody(body)
		rwb = rwb.SetHeader(httputil.ContentType, mimetype.JSON.Value())
	}
	cfg := &requestOptionConfig[T]{
		skipTokenPlacement: false,
		rwb:                rwb,
	}
	for _, option := range options {
		cfg.rwb = option.apply(cfg)
	}
	rwb = rwb.SetError(&ApiError{})
	rwb = rwb.SetResult(&ApiResult[T]{})
	if !cfg.skipTokenPlacement {
		token, err := getActiveAccountToken(ctx)
		if err != nil {
			return cfg, nil, err
		} else if token != "" {
			rwb.SetTokenAuth(token)
		}
	}
	response, err := rwb.Send()
	if err != nil {
		return cfg, nil, err
	}
	if !response.IsSuccess() {
		var apiErr ApiError
		if errors.As(response.Error(), &apiErr) {
			apiErr.Headers = response.Header()
			apiErr.StatusCode = response.StatusCode()
			return cfg, nil, apiErr
		}
		return nil, nil, response.Error()
	}
	result := response.Result()
	apiResponse, ok := result.(*ApiResult[T])
	if !ok {
		return cfg, nil, fmt.Errorf("api response result is %T not a %T", result, &ApiResult[T]{})
	}
	apiResponse.Headers = response.Header()
	apiResponse.StatusCode = response.StatusCode()
	return cfg, apiResponse, nil
}

func isNil[T any](val T) bool {
	v := any(val)
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	//nolint:exhaustive
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
