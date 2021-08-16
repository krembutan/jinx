package jinx

import (
	"context"
	"net/http"
)

var (
	CtxVersion = ctxKeyVersion{Name: "context version"}
)

type ctxKeyVersion struct {
	Name string
}

func (r *ctxKeyVersion) String() string {
	return "context value " + r.Name
}

type Meta struct {
	Code    string `json:"code,omitempty"`
	Type    string `json:"error_type,omitempty"`
	Message string `json:"error_message,omitempty"`
}

type Version struct {
	Number string `json:"number,omitempty"`
}

type Response struct {
	Version    interface{} `json:"version,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func SemanticVersion(r *http.Request, version string) {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxVersion, Version{
		Number: version,
	}))
}
