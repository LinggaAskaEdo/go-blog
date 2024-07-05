package middleware

import (
	"net/http"
	"strings"

	"golang.org/x/exp/slices"
)

func (m *middleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strMethods := []string{"GET", "POST"}

		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(strMethods, ", "))

		if !slices.Contains(strMethods, r.Method) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
