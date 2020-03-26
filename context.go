package context

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// Context represents the context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Context struct {
	echo.Context
}

var _ echo.Context = &Context{}

// RedirectHTML redirects the request to a provided URL with status code.
// By this order
// 1.HTML redirect
// 2.JavaScript redirect
func (c *Context) RedirectHTML(code int, url string) error {
	if code < 300 || code > 308 {
		return echo.ErrInvalidRedirectCode
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	c.Response().WriteHeader(code)
	html := "<html><head><meta http-equiv='Refresh' content='0; URL=" + url + "'></head><body><script>window.location.replace('" + url + "');</script></body></html>"
	_, err := c.Response().Write([]byte(html))
	return err
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

// FormValueDate returns the form field date value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date
func (c *Context) FormValueDate(name string) time.Time {
	out, err := time.Parse("2006-01-02", strings.TrimSpace(c.FormValue(name)))
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueTime returns the form field time value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
func (c *Context) FormValueTime(name string) time.Time {
	out, err := time.Parse("15:04", strings.TrimSpace(c.FormValue(name)))
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueDateTime returns the form field datetime-local value for the provided name.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
func (c *Context) FormValueDateTime(name string) time.Time {
	out, err := time.Parse("2006-01-02T15:04", strings.TrimSpace(c.FormValue(name)))
	if err != nil {
		out = time.Time{}
	}
	return out
}

// FormValueBase64 returns the form field value for the provided name.
//
// If value encoded with base64 return will be decoded string.
func (c *Context) FormValueBase64(name string) string {
	v := c.FormValueTrim(name)
	if de, err := base64.URLEncoding.DecodeString(v); err == nil {
		v = string(de)
	}
	return v
}

// FormValueInt returns the form field value for the provided name, as int.
//
// If not found returns -1 and a non-nil error.
func (c *Context) FormValueInt(name string) (int, error) {
	v := c.FormValueTrim(name)
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
	v := c.FormValueTrim(name)
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
	v := c.FormValueTrim(name)
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
	v := c.FormValueTrim(name)
	if v == "" {
		return false, echo.ErrNotFound
	}

	return strconv.ParseBool(v)
}

// ParamDefault returns path parameter by name.
//
// Returns the "def" if not found.
func (c *Context) ParamDefault(name string, def string) string {
	if v := c.Param(name); len(v) > 0 {
		return v
	}
	return def
}

// ParamTrim returns path parameter by name, without trailing spaces.
func (c *Context) ParamTrim(name string) string {
	return strings.TrimSpace(c.Param(name))
}

// ParamBase64 returns path parameter by name.
//
// If value encoded with base64 return will be decoded string.
func (c *Context) ParamBase64(name string) string {
	v := c.ParamTrim(name)
	if de, err := base64.URLEncoding.DecodeString(v); err == nil {
		v = string(de)
	}
	return v
}

// ParamInt returns path parameter by name, as int.
//
// If not found returns -1 and a non-nil error.
func (c *Context) ParamInt(name string) (int, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.Atoi(v)
}

// ParamIntDefault returns path parameter by name, as int.
//
// If not found returns or parse errors the "def".
func (c *Context) ParamIntDefault(name string, def int) int {
	if v, err := c.ParamInt(name); err == nil {
		return v
	}

	return def
}

// ParamInt64 returns path parameter by name, as float64.
//
// If not found returns -1 and a no-nil error.
func (c *Context) ParamInt64(name string) (int64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.ParseInt(v, 10, 64)
}

// ParamInt64Default returns path parameter by name, as int64.
//
// If not found or parse errors returns the "def".
func (c *Context) ParamInt64Default(name string, def int64) int64 {
	if v, err := c.ParamInt64(name); err == nil {
		return v
	}

	return def
}

// ParamFloat64 returns path parameter by name, as float64.
//
// If not found returns -1 and a non-nil error.
func (c *Context) ParamFloat64(name string) (float64, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return -1, echo.ErrNotFound
	}
	return strconv.ParseFloat(v, 64)
}

// ParamFloat64Default returns path parameter by name, as float64.
//
// If not found or parse errors returns the "def".
func (c *Context) ParamFloat64Default(name string, def float64) float64 {
	if v, err := c.ParamFloat64(name); err == nil {
		return v
	}

	return def
}

// ParamBool returns path parameter by name, as bool.
//
// If not found or value is false, then it returns false, otherwise true.
func (c *Context) ParamBool(name string) (bool, error) {
	v := c.ParamTrim(name)
	if v == "" {
		return false, echo.ErrNotFound
	}

	return strconv.ParseBool(v)
}
