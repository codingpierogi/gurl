package fetch

import (
	"fmt"
	"io"
	"net/http"
)

type Options struct {
	Verbose bool
}

func Body(url string, options Options) (string, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("user-agent", "curl/7.68.0")
	req.Header.Add("accept", "*/*")

	res, err := client.Do(req)

	if options.Verbose {
		fmt.Printf("> %s %s %s\n", req.Method, req.URL.Path, req.Proto)
		fmt.Printf("> Host: %s\n", req.Host)
		fmt.Printf("> user-agent: %s\n", req.UserAgent())
		fmt.Printf("> accept: %s\n", req.Header.Get("accept"))
		fmt.Println(">")
		fmt.Println("*")
		fmt.Printf("< %s %d\n", res.Proto, res.StatusCode)
		fmt.Printf("< date: %s\n", res.Header.Get("date"))
		fmt.Printf("< expires: %s\n", res.Header.Get("expires"))
		fmt.Printf("< cache-control: %s\n", res.Header.Get("cache-control"))
		fmt.Printf("< content-type: %s\n", res.Header.Get("content-type"))
		fmt.Println("<")
	}

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
