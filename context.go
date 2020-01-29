package context

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Context represents the context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Context struct {
	echo.Context
}

// Redirect redirects the request to a provided URL with status code.
// By this order
// 1.HTTP redirect
// 2.HTML redirect
// 3.JavaScript redirect
func (c *Context) Redirect(code int, url string) error {
	if code < 300 || code > 308 {
		return echo.ErrInvalidRedirectCode
	}
	c.Response().Header().Set(echo.HeaderLocation, url)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
	c.Response().WriteHeader(code)
	html := "<html><head><meta http-equiv='Refresh' content='0; URL=" + url + "'></head><body><script>window.location.replace('" + url + "');</script></body></html>"
	c.Response().Write([]byte(html))
	return nil
}

// FormValueDefault returns the form field value for the provided name.
//
// Returns the "def" if not found.
func (c *Context) FormValueDefault(name string, def string) string {
	if v := c.FormValue(name); len(v) > 0 {
		return v
	}
	return def
}

// FormValueTrim returns the form field value for the provided name, without trailing spaces.
func (c *Context) FormValueTrim(name string) string {
	return strings.TrimSpace(c.FormValue(name))
}

// FormValueBase64 returns the form field value for the provided name.
//
// If value encoded with base64 return will be decoded string.
func (c *Context) FormValueBase64(name string) string {
	v := c.FormValueTrim(name)
	if de, err := base64.URLEncoding.DecodeString(v); err != nil {
		v = string(de)
	}
	return v
}

// FormValueInt returns the form field value for the provided name, as int.
//
// If not found returns -1 and a non-nil error.
func (c *Context) FormValueInt(name string) (int, error) {
	v := c.FormValue(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.Atoi(v)
}

// FormValueIntDefault returns the form field value for the provided name, as int.
//
// If not found returns or parse errors the "def".
func (c *Context) FormValueIntDefault(name string, def int) int {
	if v, err := c.FormValueInt(name); err == nil {
		return v
	}

	return def
}

// FormValueInt64 returns the form field value for the provided name, as float64.
//
// If not found returns -1 and a no-nil error.
func (c *Context) FormValueInt64(name string) (int64, error) {
	v := c.FormValue(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.ParseInt(v, 10, 64)
}

// FormValueInt64Default returns the form field value for the provided name, as int64.
//
// If not found or parse errors returns the "def".
func (c *Context) FormValueInt64Default(name string, def int64) int64 {
	if v, err := c.FormValueInt64(name); err == nil {
		return v
	}

	return def
}

// FormValueFloat64 returns the form field value for the provided name, as float64.
//
// If not found returns -1 and a non-nil error.
func (c *Context) FormValueFloat64(name string) (float64, error) {
	v := c.FormValue(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.ParseFloat(v, 64)
}

// FormValueFloat64Default returns the form field value for the provided name, as float64.
//
// If not found or parse errors returns the "def".
func (c *Context) FormValueFloat64Default(name string, def float64) float64 {
	if v, err := c.FormValueFloat64(name); err == nil {
		return v
	}

	return def
}

// FormValueBool returns the form field value for the provided name, as bool.
//
// If not found or value is false, then it returns false, otherwise true.
func (c *Context) FormValueBool(name string) (bool, error) {
	v := c.FormValue(name)
	if v == "" {
		return false, echo.ErrNotFound
	}

	return strconv.ParseBool(v)
}
