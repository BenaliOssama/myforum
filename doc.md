# Documentation for Building MyForum

## Server Configuration

- If you omit the host (like we did with `:4000`), the server will listen on all your computer’s available network interfaces.

- Because DefaultServeMux is a global variable, any package can access it and register a route
— including any third-party packages that your application imports. If one of those third-
party packages is compromised, they could use DefaultServeMux to expose a malicious
handler to the web.

## header configuration
```go
// Set a new cache-control header. If an existing "Cache-Control" header exists
// it will be overwritten.
w.Header().Set("Cache-Control", "public, max-age=31536000")
// In contrast, the Add() method appends a new "Cache-Control" header and can
// be called multiple times.
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")
// Delete all values for the "Cache-Control" header.
w.Header().Del("Cache-Control")
// Retrieve the first value for the "Cache-Control" header.
w.Header().Get("Cache-Control")
// Retrieve a slice of all values for the "Cache-Control" header.
w.Header().Values("Cache-Control")
```
