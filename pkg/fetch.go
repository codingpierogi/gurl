package fetch

import (
	"fmt"
	"io"
	"net/http"
)

func Body(url string) (string, error) {
	res, err := http.Get(url)

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
	fmt.Println(res.Request.URL.Opaque)
	// for n, v := range res.Request{
	// 	fmt.Println(n, v)
	// }
	fmt.Printf("> GET %s \n", res.Request.Proto)
	fmt.Printf("> Host: %s \n", res.Request.Host)
	fmt.Printf("> user-agent: %s \n", res.Request.UserAgent())
	fmt.Printf("> accept: %s \n", res.Request.Header["Accept"])

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
