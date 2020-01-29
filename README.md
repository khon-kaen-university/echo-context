<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>

## echo-context
Add some helpful helper function for [Echo](https://github.com/labstack/echo) 4 Go web framework's Context.

### Example

```go
package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/labstack/echo/v4/middleware"
  zercleCTX "github.com/khon-kaen-university/echo-context"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Use in middleware
  e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := &zercleCTX.Context{c}
		return next(ctx)
	}
  })

  // Use in routes
  e.GET("/", func(c echo.Context) error {
    ctx := &zercleCTX.Context{c}
    name := ctx.FormValueDefault("name", "Anonymous")
	return ctx.String(http.StatusOK, "Hello " + name)
  })

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}
```