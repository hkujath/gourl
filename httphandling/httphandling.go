package httphandling

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func GetHttpRequest(input string, printHeader bool) io.ReadCloser {

	resp, err := http.Get(input)

	if err != nil {
		fmt.Printf("Error while reading URL [%s]\n%s\n", input, err)
	}

	if printHeader {
		for k, v := range resp.Header {
			fmt.Printf("%s :\n ", k)
			for i, l := range v {
				fmt.Printf(" %03d: %s\n", i+1, l)
			}
		}
	}
	return resp.Body
}

func ValidateURL(input string) bool {
	validURL := regexp.MustCompile("^http(s)?://[[:graph:]]+")
	return validURL.MatchString(input)
}
