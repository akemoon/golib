package myhttp

import "net/http"

type ResponseWriter struct {
	ResponseWriter http.ResponseWriter
	Status         int
}

func (w *ResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.Status = code
}
