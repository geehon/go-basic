package httpserver

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetUrl() {
	r, err := http.Get("https://example.com/api/v1/server")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("body=%s\n", bs)
}
