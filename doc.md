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

## database
Use the Prepare method to create a new prepared statement for the
current connection pool. This returns a sql.Stmt object which represents
the prepared statement.

## forms data

In our code above, we accessed the form values via the r.PostForm map. But an alternative
approach is to use the (subtly different) r.Form map.
The r.PostForm map is populated only for POST , PATCH and PUT requests, and contains the
form data from the request body.
In contrast, the r.Form map is populated for all requests (irrespective of their HTTP method),
and contains the form data from any request body and any query string parameters. So, if our
form was submitted to /snippet/create?foo=bar , we could also get the value of the foo
parameter by calling r.Form.Get("foo") . Note that in the event of a conflict, the request
body value will take precedent over the query string parameter.
Using the r.Form map can be useful if your application sends data in a HTML form and in the
URL, or you have an application that is agnostic about how parameters are passed. But in ourcase those things aren’t applicable. We expect our form data to be sent in the request body
only, so it’s for sensible for us to access it via r.PostForm .


## sessions
So, what happens in our application is that the LoadAndSave() middleware checks each
incoming request for a session cookie. If a session cookie is present, it reads the session token
and retrieves the corresponding session data from the database (while also checking that the
session hasn’t expired). It then adds the session data to the request context so it can be used
in your handlers.

## security
```bash
    go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
2022/02/17 18:51:29 wrote cert.pem
2022/02/17 18:51:29 wrote key.pem

## further reading
### Disable FileServer Directory Listings
https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
### spread packages
https://gist.github.com/alexedwards/5cd712192b4831058b21
### debug.stack()
https://pkg.go.dev/runtime/debug#Stack
