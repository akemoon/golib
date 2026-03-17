package httplib

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

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
