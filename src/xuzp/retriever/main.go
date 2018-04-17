package main

import (
	"fmt"
	"xuzp/retriever/mock"
	"xuzp/retriever/real"
	"time"
	"go/types"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

const url string = "http://www.imooc.com"
func post(poster Poster) {
	poster.Post(url string, map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	Connect(host string)
}

func session(s RetrieverPoster) {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com"
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is imooc.com"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")

	//fmt.Println(session(r))

	inspect(r)

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
