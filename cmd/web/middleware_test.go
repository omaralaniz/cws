package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Fake handler we can pass to the middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PING"))
	})

	//Passes the fake handler and passes the recorder and fake request
	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	//Gets the headers of the response
	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("Expected: %q Output: %q", "deny", frameOptions)
	}

	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("Expected: %q Output: %q", "1; mode=block", xssProtection)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "PING" {
		t.Errorf("Expected: %q", "PING")
	}

}
