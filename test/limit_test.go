package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkLimiter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/comments", nil)
		request.Header.Add("content-type", "application/json")
		TestRouter.ServeHTTP(response, request)
		log.Println(response.Code, response.Body)
	}
}
