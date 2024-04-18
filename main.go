package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	var headerFlags []string
	var body string

	flag.String("body", "", "Put valid JSON")
	flag.String("H", "", "")

	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		switch f.Name {
		case "H":
			headerFlags = append(headerFlags, f.Value.String())
		case "body":
			body = f.Value.String()
		}
	})

	var httpMethod = flag.Arg(0)
	var url = flag.Arg(1)

	var client = &http.Client{}
	var requestBody = strings.NewReader(body)

	var request, requestError = http.NewRequest(httpMethod, url, requestBody)

	if requestError != nil {
		fmt.Println("Error: ", requestError)
	}

	var response, responseError = client.Do(request)

	if responseError != nil {
		fmt.Println("Error: ", responseError)
	}

	fmt.Println(response.Status)
}
