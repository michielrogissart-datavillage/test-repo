// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleApiKey handles ApiKey security.
	HandleApiKey(ctx context.Context, operationName OperationName, t ApiKey) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityApiKey(ctx context.Context, operationName OperationName, req *http.Request) (context.Context, bool, error) {
	var t ApiKey
	const parameterName = "X-Cage-Id"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleApiKey(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// ApiKey provides ApiKey security value.
	ApiKey(ctx context.Context, operationName OperationName) (ApiKey, error)
}

func (s *Client) securityApiKey(ctx context.Context, operationName OperationName, req *http.Request) error {
	t, err := s.sec.ApiKey(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"ApiKey\"")
	}
	req.Header.Set("X-Cage-Id", t.APIKey)
	return nil
}
