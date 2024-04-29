package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	headers []string
	body    string
	method  string
	url     string
}

func ParseFlags() Request {
	req := Request{
		headers: []string{},
		body:    "",
		method:  "",
		url:     "",
	}

	flag.String("body", "", "request body")
	flag.String("H", "", "request header")

	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		switch f.Name {
		case "H":
			req.headers = append(req.headers, f.Value.String())
		case "body":
			req.body = f.Value.String()
		}
	})

	req.method = flag.Arg(0)
	req.url = flag.Arg(1)
	return req
}

func PerformRequest(req *Request) *http.Response {
	client := &http.Client{}
	requestBody := strings.NewReader(req.body)

	request, requestError := http.NewRequest(req.method, req.url, requestBody)

	if requestError != nil {
		fmt.Println("Error: ", requestError)
	}

	response, responseError := client.Do(request)

	if responseError != nil {
		fmt.Println("Error: ", responseError)
	}

	return response
}

func DisplayResults(res *http.Response) {
	fmt.Println(res.Status)
	fmt.Println()

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer res.Body.Close()

	fmt.Println(string(responseBody))
}

func main() {
	req := ParseFlags()
	res := PerformRequest(&req)
	DisplayResults(res)
}

