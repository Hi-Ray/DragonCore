package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Method string

const (
	GET   		Method = http.MethodGet
	POST  		Method = http.MethodPost
	PUT    		Method = http.MethodPut
	DELETE   	Method = http.MethodDelete
	PATCH    	Method = http.MethodPatch
	OPTIONS  	Method = http.MethodOptions
	CONNECT   	Method = http.MethodConnect
	TRACE  		Method = http.MethodTrace
	HEAD 		Method = http.MethodHead
)

type Controller interface {
	Middleware(mw echo.MiddlewareFunc) *controller
	Handler(method Method, path string, handle echo.HandlerFunc, mw ...echo.MiddlewareFunc) *controller
	Register(e *echo.Echo)
}

// controller struct
type controller struct {
	prefix string
	mw []echo.MiddlewareFunc
	h []handler
}

// handler struct
type handler struct {
	method Method
	path string
	h echo.HandlerFunc
	mw []echo.MiddlewareFunc
}

// New creates a controller
func New(prefix string) *controller {
	return &controller{prefix, []echo.MiddlewareFunc{}, []handler{}}
}

// Middleware adds a controller middleware
func (c *controller) Middleware(mw echo.MiddlewareFunc) *controller {
	c.mw = append(c.mw, mw)
	return c
}

// Handler adds a controller handler
func (c *controller) Handler(method Method, path string, handle echo.HandlerFunc, mw ...echo.MiddlewareFunc) *controller {
	c.h = append(c.h, handler{method, path, handle, mw})
	return c
}

// Register registers the controller
func (c *controller) Register(e *echo.Echo) {
	g := e.Group(c.prefix, c.mw...)
	for _, h := range c.h {
		g.Add(string(h.method), h.path, h.h, h.mw...)
	}
}
