package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldSuccess(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	res, err := testClient.Get(testServer.URL + "/hello")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World!", string(body))
	assert.Equal(t, 200, res.StatusCode)
}

func TestHelloWorldFailure(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	res, err := testClient.Post(testServer.URL+"/hello", "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, res.StatusCode)
}
