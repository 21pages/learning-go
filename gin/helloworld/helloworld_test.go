package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloworld(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil) //GET必须大写
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, "pong", w.Body.String())
}
