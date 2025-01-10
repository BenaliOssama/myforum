# Documentation for Building MyForum

## project structure
The cmd directory will contain the application-specific code for the executable applications
in the project. For now we’ll have just one executable application — the web application —
which will live under the cmd/web directory.

The internal directory will contain the ancillary non-application-specific code used in the
project. We’ll use it to hold potentially reusable code like validation helpers and the SQL
database models for the project.

any packages which live under this directory can only be imported by codeinside the parent of the internal directory

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
## further reading
### Disable FileServer Directory Listings
https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
### spread packages
https://gist.github.com/alexedwards/5cd712192b4831058b21
### debug.stack()
https://pkg.go.dev/runtime/debug#Stack
