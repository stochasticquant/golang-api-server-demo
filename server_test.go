package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// **************************************************************************************************************************************
// Tests for endpoint 1
// **************************************************************************************************************************************

func TestHelloWorldPass(t *testing.T) {

	// create a test server
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	// create a test client
	testClient := testServer.Client()

	// test client send GET request to server for testing
	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Fatalf("Error sending request to server: %s", err)
	}

	// returns the body as a slice of bits
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}

	// test the response status
	assert.Equal(t, 200, response.StatusCode, http.StatusOK)

	// test the response body
	assert.Equal(t, "Hello, World!", string(body))

}

func TestHelloWorldFail(t *testing.T) {

	// create a test server
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	// create a test client
	testClient := testServer.Client()

	// test client send GET request to server for testing
	fmt.Println(testServer.URL)

	body := strings.NewReader("Test body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Fatalf("Error sending request to server: %s", err)
	}

	// test the response status
	assert.Equal(t, 405, response.StatusCode, http.StatusOK)

}

// **************************************************************************************************************************************
// Tests for endpoint 2
// **************************************************************************************************************************************

func TestHealthPass(t *testing.T) {

	// create a test server
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	// create a test client
	testClient := testServer.Client()

	// test client send GET request to server for testing
	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Fatalf("Error sending request to server: %s", err)
	}

	// returns the body as a slice of bits
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}

	// test the response status
	assert.Equal(t, 200, response.StatusCode, http.StatusOK)

	// test the response body
	assert.Equal(t, "OK!", string(body))

}

func TestHealthFail(t *testing.T) {

	// create a test server
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	// create a test client
	testClient := testServer.Client()

	// test client send GET request to server for testing
	fmt.Println(testServer.URL)

	body := strings.NewReader("Test body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Fatalf("Error sending request to server: %s", err)
	}

	// test the response status
	assert.Equal(t, 405, response.StatusCode, http.StatusOK)

}
