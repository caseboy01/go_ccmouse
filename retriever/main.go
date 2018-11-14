package main

import (
	"fmt"
	"mooc_google_shizhan/retriever/mock"
	"mooc_google_shizhan/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.xingdong365.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(
		"http://xingdong365.com",
		map[string]string{
			"name":     "action",
			"language": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://xingdong365.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked xingdong365.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is fake xindong365.com"}
	//fmt.Printf("%T %v \n",r,r)

	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Printf("%T %v \n",r,r)
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println(session(&retriever))

	//fmt.Println(download(r))

}
