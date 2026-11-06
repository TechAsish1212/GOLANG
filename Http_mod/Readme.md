# HTTP in Go

This module demonstrates how to create a basic HTTP server using Go's standard
`net/http` package.

## Run

```bash
go run .
```

The server listens on `http://localhost:8080`.

## Test

```bash
curl http://localhost:8080
```

## Key concepts

- `http.HandleFunc` registers a route handler.
- `http.ListenAndServe` starts the server.
- A handler receives an `http.ResponseWriter` and `*http.Request`.
