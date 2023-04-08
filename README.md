# cors

**CORS (Cross-Origin Resource Sharing)** is a security feature implemented in web browsers that restricts web pages from making requests to a different domain than the one that served the web page. CORS middleware is used to enable cross-origin requests in web applications.

This package provides a middleware that can be used to enable CORS in a `Lightning` go web application. The middleware can be configured with various options to allow specific origins, HTTP methods, headers, and more.

## Installation

```bash
go get github.com/lightning-contrib/cors
```

## Usage

To use the cors middleware in your Go web application, you can import the cors package and add the middleware to your middleware chain. Here's an example:

```go
package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/cors"
)

func main() {
	app := lightning.DefaultApp()

    // Add the CORS middleware to the middleware chain
	app.Use(cors.Default())

    // Add your routes here
	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.Success("hello world")
	})

    // Start the server
	app.Run()
}
```

By default, the cors middleware allows all origins (*), all HTTP methods (GET, POST, PUT, DELETE), all headers (*), and exposes all headers (*). You can customize these options by using the various option functions provided by the cors package.

Here's an example of how to customize the allowed origins:

```go
app.Use(cors.New(
    cors.AllowOrigin([]string{"https://example.com", "https://www.example.com"}),
))
```

This will allow requests from https://example.com and https://www.example.com.

## Options

The cors middleware can be customized with the following options:

- AllowOrigin([]string): Sets the allowed origins. By default, all origins are allowed (*).
- AllowMethods([]string): Sets the allowed HTTP methods. By default, GET, POST, PUT, and DELETE are allowed.
- AllowHeaders([]string): Sets the allowed headers. By default, all headers are allowed (*).
- ExposeHeaders([]string): Sets the exposed headers. By default, all headers are exposed (*).
- SetMaxAge(int): Sets the max age in seconds for preflight requests. By default, the max age is set to 3600 seconds (1 hour).
- AllowCredentials(bool): Sets whether credentials are allowed. By default, credentials are allowed.



## API Documentation

For detailed API documentation and usage examples, please refer to the [documentation](https://pkg.go.dev/github.com/lightning-contrib/cors).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](https://github.com/lightning-contrib/cors/blob/main/CONTRIBUTING.md) for more information.

## License

This middleware is licensed under the [MIT License](https://github.com/lightning-contrib/cors/blob/main/LICENSE). 
