package fetch

import (
	"fmt"
	"io"
	"net/http"
)

func Body(url string) (string, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, )

	req.Header.Add("user-agent", "curl/7.68.0")
	
	res, err := client.Get(url)

	// > GET / HTTP/2
	// > Host : www.google.com
	// > user-agent: curl/7.68.0
	// > accept: */*
	// >
	// *
	// < HTTP/2 200
	// < date: Sat, 17 Dec 2022 14:41:44 GMT
	// < expires: -1
	// < cache-control: private, max-age=0
	// < content-type: text/html; charset=ISO-8859-1
	// <

	fmt.Println(res.Request.URL.EscapedPath())

	fmt.Printf("> GET %s %s\n", res.Request.URL.Path, res.Request.Proto)
	fmt.Printf("> Host: %s\n", res.Request.Host)
	fmt.Printf("> user-agent: %s\n", res.Request.UserAgent())
	fmt.Printf("> accept: %s\n", res.Request.Header["Accept"])
	fmt.Println(">")
	fmt.Printf("< %s %d\n", res.Proto, res.StatusCode)
	fmt.Printf("< date: %s\n", res.Header.Get("date"))
	fmt.Printf("< expires: %s\n", res.Header.Get("expires"))
	fmt.Printf("< cache-control: %s\n", res.Header.Get("cache-control"))
	fmt.Printf("< content-type: %s\n", res.Header.Get("content-type"))
	fmt.Println("<")

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
