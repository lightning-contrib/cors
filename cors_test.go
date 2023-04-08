package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-labx/lightning"
)

func TestCORS_With_No_Origin(t *testing.T) {
	// Create a new lightning app
	app := lightning.NewApp()

	// Add the CORS middleware to the middleware chain
	app.Use(Default())

	// Create a new test request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a new test response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	app.ServeHTTP(rec, req)

	// Check the Access-Control-Allow-Origin header
	if rec.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Origin"))
	}

	// Check the Access-Control-Allow-Credentials header
	if rec.Header().Get("Access-Control-Allow-Credentials") != "" {
		t.Errorf("Expected Access-Control-Allow-Credentials header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Credentials"))
	}

	// Check the Access-Control-Expose-Headers header
	if rec.Header().Get("Access-Control-Expose-Headers") != "" {
		t.Errorf("Expected Access-Control-Expose-Headers header to be empty, but got '%s'", rec.Header().Get("Access-Control-Expose-Headers"))
	}

	// Check the Access-Control-Allow-Methods header
	if rec.Header().Get("Access-Control-Allow-Methods") != "" {
		t.Errorf("Expected Access-Control-Allow-Methods header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Methods"))
	}

	// Check the Access-Control-Allow-Headers header
	if rec.Header().Get("Access-Control-Allow-Headers") != "" {
		t.Errorf("Expected Access-Control-Allow-Headers header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Headers"))
	}

	// Check the Access-Control-Max-Age header
	if rec.Header().Get("Access-Control-Max-Age") != "" {
		t.Errorf("Expected Access-Control-Max-Age header to be empty, but got '%s'", rec.Header().Get("Access-Control-Max-Age"))
	}

	// Check the status code
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status code to be %d, but got %d", http.StatusNotFound, rec.Code)
	}
}

func TestCORS(t *testing.T) {
	// Create a new lightning app
	app := lightning.NewApp()

	// Add the CORS middleware to the middleware chain
	app.Use(Default())

	// Create a new test request
	req := httptest.NewRequest("OPTIONS", "/", nil)

	// Set the Origin header
	req.Header.Set("Origin", "https://example.com")

	// Create a new test response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	app.ServeHTTP(rec, req)

	// Check the Access-Control-Allow-Origin header
	if rec.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be 'https://example.com', but got '%s'", rec.Header().Get("Access-Control-Allow-Origin"))
	}

	// Check the Access-Control-Allow-Credentials header
	if rec.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials header to be 'true', but got '%s'", rec.Header().Get("Access-Control-Allow-Credentials"))
	}

	// Check the Access-Control-Expose-Headers header
	if rec.Header().Get("Access-Control-Expose-Headers") != "*" {
		t.Errorf("Expected Access-Control-Expose-Headers header to be '*', but got '%s'", rec.Header().Get("Access-Control-Expose-Headers"))
	}

	// Check the Access-Control-Allow-Methods header
	if rec.Header().Get("Access-Control-Allow-Methods") != "GET,POST,PUT,DELETE" {
		t.Errorf("Expected Access-Control-Allow-Methods header to be 'GET,POST,PUT,DELETE', but got '%s'", rec.Header().Get("Access-Control-Allow-Methods"))
	}

	// Check the Access-Control-Allow-Headers header
	if rec.Header().Get("Access-Control-Allow-Headers") != "*" {
		t.Errorf("Expected Access-Control-Allow-Headers header to be '*', but got '%s'", rec.Header().Get("Access-Control-Allow-Headers"))
	}

	// Check the Access-Control-Max-Age header
	if rec.Header().Get("Access-Control-Max-Age") != "3600" {
		t.Errorf("Expected Access-Control-Max-Age header to be '3600', but got '%s'", rec.Header().Get("Access-Control-Max-Age"))
	}

	// Check the status code
	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status code to be %d, but got %d", http.StatusNoContent, rec.Code)
	}
}

func TestCORS_GET(t *testing.T) {
	// Create a new lightning app
	app := lightning.NewApp()

	// Add the CORS middleware to the middleware chain
	app.Use(Default())

	// Create a new test request
	req := httptest.NewRequest("GET", "/", nil)

	// Set the Origin header
	req.Header.Set("Origin", "https://example.com")

	// Create a new test response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	app.ServeHTTP(rec, req)

	// Check the Access-Control-Allow-Origin header
	if rec.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be 'https://example.com', but got '%s'", rec.Header().Get("Access-Control-Allow-Origin"))
	}

	// Check the Access-Control-Allow-Credentials header
	if rec.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials header to be 'true', but got '%s'", rec.Header().Get("Access-Control-Allow-Credentials"))
	}

	// Check the Access-Control-Expose-Headers header
	if rec.Header().Get("Access-Control-Expose-Headers") != "*" {
		t.Errorf("Expected Access-Control-Expose-Headers header to be '*', but got '%s'", rec.Header().Get("Access-Control-Expose-Headers"))
	}

	// Check the Access-Control-Allow-Methods header
	if rec.Header().Get("Access-Control-Allow-Methods") != "" {
		t.Errorf("Expected Access-Control-Allow-Methods header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Methods"))
	}

	// Check the Access-Control-Allow-Headers header
	if rec.Header().Get("Access-Control-Allow-Headers") != "" {
		t.Errorf("Expected Access-Control-Allow-Headers header to be empty, but got '%s'", rec.Header().Get("Access-Control-Allow-Headers"))
	}

	// Check the Access-Control-Max-Age header
	if rec.Header().Get("Access-Control-Max-Age") != "" {
		t.Errorf("Expected Access-Control-Max-Age header to be empty, but got '%s'", rec.Header().Get("Access-Control-Max-Age"))
	}

	// Check the status code
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status code to be %d, but got %d", http.StatusNotFound, rec.Code)
	}
}

func TestCORS_With_Options(t *testing.T) {
	// Create a new lightning app
	app := lightning.NewApp()

	// Add the CORS middleware to the middleware chain
	app.Use(New(
		AllowOrigin([]string{"https://example.com"}),
		AllowMethods([]string{"GET", "POST"}),
		AllowHeaders([]string{"Foo", "Bar"}),
		ExposeHeaders([]string{"Foo", "Bar"}),
		SetMaxAge(1000),
		AllowCredentials(true),
	))

	// Create a new test request
	req := httptest.NewRequest("OPTIONS", "/", nil)

	// Set the Origin header
	req.Header.Set("Origin", "https://example.com")

	// Create a new test response recorder
	rec := httptest.NewRecorder()

	// Perform the request
	app.ServeHTTP(rec, req)

	// Check the Access-Control-Allow-Origin header
	if rec.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be 'https://example.com', but got '%s'", rec.Header().Get("Access-Control-Allow-Origin"))
	}

	// Check the Access-Control-Allow-Credentials header
	if rec.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials header to be 'true', but got '%s'", rec.Header().Get("Access-Control-Allow-Credentials"))
	}

	// Check the Access-Control-Expose-Headers header
	if rec.Header().Get("Access-Control-Expose-Headers") != "Foo,Bar" {
		t.Errorf("Expected Access-Control-Expose-Headers header to be 'Foo,Bar', but got '%s'", rec.Header().Get("Access-Control-Expose-Headers"))
	}

	// Check the Access-Control-Allow-Methods header
	if rec.Header().Get("Access-Control-Allow-Methods") != "GET,POST" {
		t.Errorf("Expected Access-Control-Allow-Methods header to be 'GET,POST', but got '%s'", rec.Header().Get("Access-Control-Allow-Methods"))
	}

	// Check the Access-Control-Allow-Headers header
	if rec.Header().Get("Access-Control-Allow-Headers") != "Foo,Bar" {
		t.Errorf("Expected Access-Control-Allow-Headers header to be 'Foo,Bar', but got '%s'", rec.Header().Get("Access-Control-Allow-Headers"))
	}

	// Check the Access-Control-Max-Age header
	if rec.Header().Get("Access-Control-Max-Age") != "1000" {
		t.Errorf("Expected Access-Control-Max-Age header to be '1000', but got '%s'", rec.Header().Get("Access-Control-Max-Age"))
	}

	// Check the status code
	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status code to be %d, but got %d", http.StatusNoContent, rec.Code)
	}
}
