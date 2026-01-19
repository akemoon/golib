package middleware

import "net/http"

type Midddleware func(http.Handler) http.Handler
