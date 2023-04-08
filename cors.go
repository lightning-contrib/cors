package cors

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-labx/lightning"
)

// config struct holds the configuration for CORS middleware
type config struct {
	allowedOrigins   []string // allowed origins
	allowedMethods   []string // allowed HTTP methods
	allowedHeaders   []string // allowed headers
	exposedHeaders   []string // exposed headers
	maxAge           int      // max age
	allowCredentials bool     // allow credentials
}

// option is a function that takes a pointer to a config struct and modifies it
type option func(*config)

// AllowOrigin is an option function that sets the allowed origins
func AllowOrigin(origin []string) option {
	return func(cfg *config) {
		cfg.allowedOrigins = origin
	}
}

// AllowMethods is an option function that sets the allowed HTTP methods
func AllowMethods(methods []string) option {
	return func(cfg *config) {
		cfg.allowedMethods = methods
	}
}

// AllowHeaders is an option function that sets the allowed headers
func AllowHeaders(headers []string) option {
	return func(cfg *config) {
		cfg.allowedHeaders = headers
	}
}

// ExposeHeaders is an option function that sets the exposed headers
func ExposeHeaders(headers []string) option {
	return func(cfg *config) {
		cfg.exposedHeaders = headers
	}
}

// SetMaxAge is an option function that sets the max age
func SetMaxAge(seconds int) option {
	return func(cfg *config) {
		cfg.maxAge = seconds
	}
}

// AllowCredentials is an option function that sets whether credentials are allowed
func AllowCredentials(allowCredentials bool) option {
	return func(cfg *config) {
		cfg.allowCredentials = allowCredentials
	}
}

// Default returns the default CORS middleware
func Default() lightning.Middleware {
	return New()
}

// New returns a new CORS middleware with the given options
func New(options ...option) lightning.Middleware {
	cfg := &config{
		allowedOrigins:   []string{"*"},
		allowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		allowedHeaders:   []string{"*"},
		exposedHeaders:   []string{"*"},
		maxAge:           3600,
		allowCredentials: true,
	}

	for _, option := range options {
		option(cfg)
	}

	return func(ctx *lightning.Context) {
		origin := ctx.Header("Origin")

		if origin == "" {
			ctx.Next()
			return
		}

		allowedOrigins := cfg.allowedOrigins

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == "*" || origin == allowedOrigin {
				ctx.SetHeader("Access-Control-Allow-Origin", origin)
				ctx.SetHeader("Access-Control-Expose-Headers", strings.Join(cfg.exposedHeaders, ","))
				ctx.SetHeader("Access-Control-Allow-Credentials", strconv.FormatBool(cfg.allowCredentials))
				break
			}
		}

		if ctx.Method == "OPTIONS" {
			ctx.SetHeader("Access-Control-Allow-Methods", strings.Join(cfg.allowedMethods, ","))
			ctx.SetHeader("Access-Control-Allow-Headers", strings.Join(cfg.allowedHeaders, ","))
			ctx.SetHeader("Access-Control-Max-Age", strconv.Itoa(cfg.maxAge))
			ctx.SetStatus(http.StatusNoContent)
		} else {
			ctx.Next()
		}
	}
}
