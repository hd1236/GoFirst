package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

/**
并没有显示指定实现了哪个接口，而只需要定义一个作为接受者的方法就会自动关联上
*/
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}

//func (r Retriever) Get(url string)string{
//
//    resp,err := http.Get(url)
//    if err != nil{
//        panic(err)
//    }
//
//    result, err := httputil.DumpRequest(resp,true)
//
//    resp.Body.Close()
//
//    if err != nil{
//        panic(err)
//    }
//
//    return string(result)
//}
