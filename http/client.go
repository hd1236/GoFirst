package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, e := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Mobile Safari/537.36")
	response, e := http.DefaultClient.Do(request)

	if e != nil {
		panic(e)
	}

	defer response.Body.Close()

	s, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)

	//查看api
	//godoc -http :8888
	//标准库中文版：https://studygolang.com/pkgdoc
}
