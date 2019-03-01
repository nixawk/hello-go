package main

import "net/url"
import "fmt"

func main() {
        s := "postgres://user:pass@host.com:5432/path?k=v#f"
        u, err := url.Parse(s)
        if err != nil {
                panic(err)
        }

        fmt.Printf("Scheme: %s\n", u.Scheme)
        fmt.Printf("Opaque: %s\n", u.Opaque)
        fmt.Printf("Host: %s\n", u.Host)
        fmt.Printf("Port: %s\n", u.Port())
        fmt.Printf("RawPath: %s\n", u.RawPath)
        fmt.Printf("RawQuery: %s\n", u.RawQuery)
        fmt.Printf("Fragment: %s\n", u.Fragment)
}

/*

type URL struct {
        Scheme     string
        Opaque     string    // encoded opaque data
        User       *Userinfo // username and password information
        Host       string    // host or host:port
        Path       string    // path (relative paths may omit leading slash)
        RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
        ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
        RawQuery   string    // encoded query values, without '?'
        Fragment   string    // fragment for references, without '#'
}

*/

/* https://golang.org/pkg/net/url/ */