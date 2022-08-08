package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getFixture(name string) string {
	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(mydir, "test_support", name)
}

func validateFixture(filename string) (string, int) {
	router := setupRouter()
	filePath := getFixture(filename)
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	w, err := mw.CreateFormFile("file", filePath)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, file); err != nil {
		panic(err)
	}
	mw.Close() // close the writer before making the request
	req, _ := http.NewRequest(http.MethodPost, "/ruby/validate", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res.Body.String(), res.Code
}

func TestRubyValidateComplex(t *testing.T) {
	bodyString, code := validateFixture("complex.rb")
	expectString := `[{"name":"foo","required":true},{"name":"a","required":false},{"name":"some_arr","required":false},{"name":"bar","required":false},{"name":"baz","required":false},{"name":"dry_run:","required":false},{"name":"other_thing:","required":false},{"name":"hi:","required":false}]`
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, expectString+"\n", bodyString)
}

func TestRubyValidateCantParse(t *testing.T) {
	bodyString, code := validateFixture("cant_parse.rb")
	expectString := `{"error":"no top-level run function detected in your program"}`
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, expectString+"\n", bodyString)
}

func TestRubyValidateSimple(t *testing.T) {
	bodyString, code := validateFixture("simple.rb")
	expectString := `[]`
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, expectString+"\n", bodyString)
}

func TestRubyValidateMissingRun(t *testing.T) {
	bodyString, code := validateFixture("missing_run.rb")
	expectString := `{"error":"no top-level run function detected in your program"}`
	assert.Equal(t, http.StatusBadRequest, code)
	assert.Equal(t, expectString+"\n", bodyString)
}
