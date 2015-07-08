package main

import (
    "fmt"
    "os"
    "net/http"
    "net/http/cookiejar"
)

type LogRedirects struct {
        Transport http.RoundTripper
}

func (l LogRedirects) RoundTrip(req *http.Request) (resp *http.Response, err error) {
    t := l.Transport
    if t == nil {
        t = http.DefaultTransport
    }
    resp, err = t.RoundTrip(req)
    if err != nil {
        return
    }
    switch resp.StatusCode {
        case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
            fmt.Println("#### Redirected ####\n")
            fmt.Println("Request for:", req.URL)
            fmt.Println("Redirected to:", resp.Header.Get("location"))
            fmt.Println("Status code:", resp.StatusCode)

            fmt.Println("\n\n#### Headers ####\n")

            for k, v := range resp.Header {
                fmt.Println(k, ": ", v[0])
            }

            fmt.Println("\n\n#### Cookies ####\n")
            for _, cookie := range resp.Header[http.CanonicalHeaderKey("Set-Cookie")] {
                fmt.Println(cookie)
            }

            if resp.Header[http.CanonicalHeaderKey("Set-Cookie")] == nil {
                fmt.Println("No cookies created.")
            }

            fmt.Println("\n")
    }
    return
}

func main() {
    // create a cookiejar
    cookieJar, _ := cookiejar.New(nil)

    // create a client
    client := http.Client{Jar: cookieJar, Transport: LogRedirects{}}

    fmt.Println("Initial request to:", os.Args[1], "\n")

    // do the request
    resp, err := client.Get(os.Args[1])
    if err != nil {
        fmt.Println("[!] Error while accessing the host.")
        return
    }
    fmt.Println("Final URL is", resp.Request.URL, " with status code: ", resp.StatusCode)

    // display information for the last request
    fmt.Println("\n\n#### Headers ####\n")
    for k, v := range resp.Header {
        fmt.Println(k, ": ", v[0])
    }

    fmt.Println("\n\n#### Cookies ####\n")
    for _, cookie := range resp.Header[http.CanonicalHeaderKey("Set-Cookie")] {
        fmt.Println(cookie)
    }

    if resp.Header[http.CanonicalHeaderKey("Set-Cookie")] == nil {
        fmt.Println("No cookies created.")
    }
}
