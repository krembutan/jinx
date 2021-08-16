package jinx

import (
	"context"
	"net/http"
)

type ctxError struct {
	Name string
}

func (r *ctxError) String() string {
	return "context value " + r.Name
}

var CtxError = ctxError{Name: "context error"}

// ErrBadRequest error http StatusBadRequest
func ErrBadRequest(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusBadRequest))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	return err
}

// ErrUnauthorized error http StatusUnauthorized
func ErrUnauthorized(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusUnauthorized))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusUnauthorized)
	return err
}

// ErrForbidden error http StatusForbidden
func ErrForbidden(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusForbidden))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusForbidden)
	return err
}

// ErrMethodNotAllowed error http StatusMethodNotAllowed
func ErrMethodNotAllowed(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusMethodNotAllowed))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusMethodNotAllowed)
	return err
}

// ErrInternalServerError error http StatusInternalServerError
func ErrInternalServerError(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusInternalServerError))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	return err
}

// ErrBadGateway error http StatusBadGateway
func ErrBadGateway(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusBadGateway))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadGateway)
	return err
}

// ErrServiceUnavailable error http StatusServiceUnavailable
func ErrServiceUnavailable(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxError, http.StatusServiceUnavailable))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusServiceUnavailable)
	return err
}
