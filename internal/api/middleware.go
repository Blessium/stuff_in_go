package api

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.Handler, m ...Middleware) http.HandlerFunc {
	wrapped := h.ServeHTTP

	if len(m) < 1 {
		return wrapped
	}

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)
	})
}
