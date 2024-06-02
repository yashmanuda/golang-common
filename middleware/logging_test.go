package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	middleware := LoggingMiddleware(sampleHandler)
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	middleware.ServeHTTP(recorder, req)
	responseBody := recorder.Body.String()
	assert.Equal(t, "OK", responseBody, "Expected response body to be 'OK'")
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200")
}
