# go-logx

[![Documentation](https://godoc.org/github.com/mashiike/go-logx?status.svg)](https://godoc.org/github.com/mashiike/go-logx)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/mashiike/go-logx)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/mashiike/go-logx)
![Github Actions test](https://github.com/mashiike/go-logx/workflows/Test/badge.svg?branch=main)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mashiike/go-logx/blob/master/LICENSE)


This package extends Go's standard package `log` to context.Context.

## Usage 

```go
logger := log.New(&buf, "context_logger: ", log.Lshortfile)
ctx := logx.WithLogger(context.Background(), logger)
logx.Print(ctx, "Hello World!?")
```

for example: http logging middleware

```go 
func main() {
	middleware := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get("X-Request-ID")
			if reqID == "" {
				reqID = "-"
			}
			logger := log.New(os.Stderr, " request_id:"+reqID, log.LUTC)
			h.ServeHTTP(w, r.WithContext(
				logx.WithLogger(r.Context(), logger),
			))
		})
	}
	http.ListenAndServe(":8080", middleware(http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request){
	logx.Print(r.Context(), "call handler")
	w.WriteHeader(http.StatusOK)
}
```

## Benchmark

```shell
$ go test -bench . -benchmem -benchtime 10s 
goos: darwin
goarch: amd64
pkg: github.com/mashiike/go-logx
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkLogDefaultLogger-8     49358439               231.9 ns/op            16 B/op          1 allocs/op
BenchmarkLogxDefaultLogger-8    48866178               276.3 ns/op            32 B/op          2 allocs/op
BenchmarkLogxContextLogger-8    47944270               277.2 ns/op            32 B/op          2 allocs/op
```

