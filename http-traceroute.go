package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type logRedirects struct {
	Transport http.RoundTripper
}

func (l logRedirects) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport
	if t == nil {
		t = http.DefaultTransport
	}

	if resp, err = t.RoundTrip(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		fmt.Println("#### Redirected ####\n")
		fmt.Println("Request for:", req.URL)
		fmt.Println("Redirected to:", resp.Header.Get("location"))
		fmt.Println("Status code:", resp.StatusCode)

		fmt.Println("\n\n#### Headers ####\n")

		var max int
		for k := range resp.Header {
			if len(k) > max {
				max = len(k)
			}
		}
		formatHeaders := fmt.Sprint("%-", max, "s: %s\n")

		for k, v := range resp.Header {
			fmt.Printf(formatHeaders, k, v[0])
		}

		fmt.Println("\n\n#### Cookies ####\n")
		if resp.Header[http.CanonicalHeaderKey("Set-Cookie")] == nil {
			fmt.Println("No cookies created.")
		} else {
			for _, cookie := range resp.Header[http.CanonicalHeaderKey("Set-Cookie")] {
				fmt.Println(cookie)
			}
		}
		fmt.Println("\n")
	}
	return
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("\nYou need to specify an URL as parameter\n")
		os.Exit(1)
	}

	// create a cookiejar
	cookieJar, _ := cookiejar.New(nil)

	// create a client
	client := http.Client{Jar: cookieJar, Transport: logRedirects{}}

	fmt.Println("Initial request to:", os.Args[1], "\n")

	// do the request
	resp, err := client.Get(os.Args[1])
	if err != nil {
		fmt.Printf("[!] Error while accessing the host: %s", err)
		return
	}
	fmt.Println("Final URL is", resp.Request.URL, " with status code: ", resp.StatusCode)

	// display information for the last request
	fmt.Println("\n\n#### Headers ####\n")

	var max int
	for k := range resp.Header {
		if len(k) > max {
			max = len(k)
		}
	}
	formatHeaders := fmt.Sprint("%-", max, "s: %s\n")

	for k, v := range resp.Header {
		fmt.Printf(formatHeaders, k, v[0])
	}

	fmt.Println("\n\n#### Cookies ####\n")

	if resp.Header[http.CanonicalHeaderKey("Set-Cookie")] == nil {
		fmt.Println("No cookies created.")
	} else {
		for _, cookie := range resp.Header[http.CanonicalHeaderKey("Set-Cookie")] {
			fmt.Println(cookie)
		}
	}
}
