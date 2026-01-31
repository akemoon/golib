package myhttp

import "net/http"

type Midddleware func(http.Handler) http.Handler
