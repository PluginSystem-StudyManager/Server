package utils

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc MiddlewareFunction) MiddlewareFunction

type MiddlewareFunction func(writer ResponseWriter, request *http.Request)

func Chain(first Middleware, others ...Middleware) Middleware {
	return func(next MiddlewareFunction) MiddlewareFunction {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return first(next)
	}
}

func Logging() Middleware {
	return func(next MiddlewareFunction) MiddlewareFunction {
		return func(writer ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				duration := time.Since(start)
				log.Printf("[%s] %s %s %d", request.Method, request.URL.Path, duration, writer.statusCode)
			}()
			writer.WriteHeader(10)
			next(writer, request)
		}
	}
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewWrapperResponseWriter(w http.ResponseWriter) ResponseWriter {
	return ResponseWriter{
		w, http.StatusOK,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type Param func(int) int

func d1() Param {
	return func(a int) int {
		return a + 1
	}
}

func d2() Param {
	return func(a int) int {
		return a + 2
	}
}

func d3() Param {
	return func(a int) int {
		return a + 3
	}
}

func setup() {
	func(f ...Param) {
		for _, n := range f {
			n(1)
		}
	}(d1(), d2(), d3())
}
