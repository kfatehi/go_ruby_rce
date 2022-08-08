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

func TestRubyValidateProper(t *testing.T) {
	router := setupRouter()
	filePath := getFixture("proper.rb")
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile("file", filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	// close the writer before making the request
	mw.Close()

	req, _ := http.NewRequest(http.MethodPost, "/ruby/validate", body)

	req.Header.Add("Content-Type", mw.FormDataContentType())

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `[{"name":"foo","required":true},{"name":"a","required":false},{"name":"some_arr","required":false},{"name":"bar","required":false},{"name":"baz","required":false},{"name":"dry_run:","required":false},{"name":"other_thing:","required":false},{"name":"hi:","required":false}]`+"\n", res.Body.String())
}
