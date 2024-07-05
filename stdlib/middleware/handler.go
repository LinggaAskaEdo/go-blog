package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
)

func (m *middleware) Handler(next http.Handler) http.Handler {
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	start := time.Now()

	// 	lrw := responseHandler(w)

	// 	defer func() {
	// 		panicVal := recover()
	// 		if panicVal != nil {
	// 			lrw.statusCode = http.StatusInternalServerError // ensure that the status code is updated
	// 			panic(panicVal)                                 // continue panicking
	// 		}

	// 		m.log.
	// 			Info().
	// 			Str("event", "START").
	// 			Str("method", r.Method).
	// 			Str("url", r.URL.RequestURI()).
	// 			Str("user_agent", r.UserAgent()).
	// 			Dur("elapsed_ms", time.Since(start)).
	// 			Int("status_code", lrw.statusCode).
	// 			Send()
	// 	}()

	// 	next.ServeHTTP(w, r)
	// })

	h := hlog.NewHandler(m.log)

	accessHandler := hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			hlog.
				FromRequest(r).
				Info().
				Str("event", "START").
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status_code", status).
				Int("response_size_bytes", size).
				Dur("elapsed_ms", duration).
				Send()
		},
	)

	userAgentHandler := hlog.UserAgentHandler("http_user_agent")

	return h(accessHandler(userAgentHandler(next)))
}
