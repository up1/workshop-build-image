package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	err := helloHandler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// Parse and verify JSON response
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", response["message"])
}

func TestHelloHandlerWithEchoServer(t *testing.T) {
	// Setup Echo server
	e := echo.New()
	setupRoutes(e)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Serve request
	e.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// Verify response body
	expectedJSON := `{"message":"Hello, World!"}`
	actualJSON := strings.TrimSpace(rec.Body.String())
	assert.JSONEq(t, expectedJSON, actualJSON)
}

func TestHelloHandlerHTTPServer(t *testing.T) {
	// Setup Echo server
	e := echo.New()
	setupRoutes(e)

	// Create test server
	server := httptest.NewServer(e)
	defer server.Close()

	// Make HTTP request to test server
	resp, err := http.Get(server.URL + "/")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	// Parse JSON response
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", response["message"])
}

func TestInvalidEndpoint(t *testing.T) {
	// Setup Echo server
	e := echo.New()
	setupRoutes(e)

	// Test invalid endpoint
	req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Should return 404
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestWrongHTTPMethod(t *testing.T) {
	// Setup Echo server
	e := echo.New()
	setupRoutes(e)

	// Test POST method on GET endpoint
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Should return 405 Method Not Allowed
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

// Benchmark test for the handler
func BenchmarkHelloHandler(b *testing.B) {
	e := echo.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		_ = helloHandler(c)
	}
}

// Example test showing how to test with custom headers
func TestHelloHandlerWithCustomHeaders(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Custom-Header", "test-value")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	err := helloHandler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify that we can access custom headers in the context if needed
	assert.Equal(t, "test-value", c.Request().Header.Get("X-Custom-Header"))
}
