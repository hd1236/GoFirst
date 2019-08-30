package main

import (
	"HelloWorld/retriever/mock"
	"HelloWorld/retriever/real"
	"fmt"
	"time"
)

/**
接口
定义一个接口，实现者不需要显示指定，只需要实现方法就会默认实现了。
real.Retriever
mock.Retriever
*/
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	ret := s.Post(url, map[string]string{
		"name": "another contents",
	})
	return s.Get(ret)
}

//使用者定义接口
const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post("http://www.baidu.com",
		map[string]string{
			"name": "hand",
			"id":   "123",
		})
}

//type switch
func inspect(r Retriever) {
	fmt.Printf("%T,%v\n", r, r)
	//接口实现类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
func main() {
	var r Retriever
	r = &mock.Retriever{"handsome boy"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//通过Type assertion或者type switch方式获取接口信息
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Print("not mockR")
	}

	realRetriever := r.(*real.Retriever)
	fmt.Print(realRetriever.TimeOut)

	//fmt.Println(download(r))

	fmt.Println("try a session")

	fmt.Println(session(&mock.Retriever{Contents: "init content"}))

}
