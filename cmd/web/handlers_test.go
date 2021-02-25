package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApp(t)
	ts := newTestServer(t, app.routes())

	defer ts.Close()
	// //The recorder records the http respsonse
	// rr := httptest.NewRecorder()

	// //Make a bogus request
	// r, err := http.NewRequest(http.MethodGet, "/", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	statusCode, _, body := ts.get(t, "/ping")

	if statusCode != http.StatusOK {
		t.Errorf("Expected: %d Output: %d", http.StatusOK, statusCode)
	}

	// //Instead of passing a normal response we pass the response recorder, and pass the bogus request
	// ping(rr, r)

	// //Result() gets the http.Response
	// rs := rr.Result()

	//Closes the Body once done

	if string(body) != "PINGED" {
		t.Errorf("Expected: %q", "PINGED")
	}
}

func TestShowArticle(t *testing.T) {
	app := newTestApp(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   []byte
	}{
		{"Valid ID", "/post/1", http.StatusOK, []byte("This mock info so pretend this is informative content")},
		{"Invalid ID", "/post/2", http.StatusNotFound, nil},
		{"Negative ID", "/post/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/post/1.23", http.StatusNotFound, nil},
		{"String ID", "/post/foo", http.StatusNotFound, nil},
		{"Empty ID", "/post/", http.StatusNotFound, nil},
		{"Trailing slash", "/post/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.url)

			if code != tt.expectedStatus {
				t.Errorf("want %d; got %d", tt.expectedStatus, code)
			}

			if !bytes.Contains(body, tt.expectedBody) {
				t.Errorf("want body to contain %q", tt.expectedBody)
			}
		})
	}

}

func TestSignupAuthor(t *testing.T) {
	app := newTestApp(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/author/signup")
	csrfToken := extractToken(t, body)

	authors := []struct {
		name           string
		userName       string
		userEmail      string
		userPassword   string
		csrfToken      string
		expectedStatus int
		expectedBody   []byte
	}{
		{"Valid submission", "Mockaton", "mockaton@gmail.com", "validPass123", csrfToken, http.StatusSeeOther, nil},
		{"Empty name", "", "mockaton@gmail.com", "validPass123", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty email", "Mockaton", "", "validPass123", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty password", "Mockaton", "mockaton@gmail.com", "", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Invalid email (incomplete domain)", "Mockaton", "mockaton@gmail.", "validPass123", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", "Mockaton", "mockatongmail.com", "validPass123", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing local part)", "Mockaton", "@gmail.com", "validPass123", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", "Mockaton", "Mockaton@gmail.com", "pass", csrfToken, http.StatusOK, []byte("This field is too short (minimum is 10 characters)")},
		{"Duplicate email", "Mockaton", "Mockaton@gmail.com", "validPass123", csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "", "bogusToken", http.StatusBadRequest, nil},
	}
	for _, aa := range authors {
		t.Run(aa.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", aa.userName)
			form.Add("email", aa.userEmail)
			form.Add("password", aa.userPassword)
			form.Add("csrf_token", aa.csrfToken)

			status, _, body := ts.postForm(t, "/author/signup", form)
			if status != aa.expectedStatus {
				t.Errorf("Expected: %d; Output: %d", aa.expectedStatus, status)
			}

			if !bytes.Contains(body, aa.expectedBody) {
				t.Errorf("Expected: %q", aa.expectedBody)
			}
		})
	}
}
