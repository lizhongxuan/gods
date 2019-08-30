package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//testV1()
	//test2()
	test3()
}
func testV1() {
	future := RequestFuture("https://api.github.com/users/octocat/orgs")
	body := <-future
	fmt.Printf("reponse length: %d", len(body))
}
func RequestFuture(url string) <-chan []byte {
	c := make(chan []byte, 1)
	go func() {
		var body []byte
		defer func() {
			c <- body
		}()

		res, err := http.Get(url)
		if err != nil {
			return
		}
		defer res.Body.Close()

		body, _ = ioutil.ReadAll(res.Body)
	}()

	return c
}

func test2() {
	futureV2 := RequestFutureV2("https://api.github.com/users/octocat/orgs")

	// not block
	fmt.Println("V2 is this locked again")

	bodyV2, err := futureV2() // block
	if err == nil {
		fmt.Printf("V2 response length %d\n", len(bodyV2))
	} else {
		fmt.Printf("V2 error is %v\n", err)
	}
}

// RequestFutureV2 return value and error
func RequestFutureV2(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		var res *http.Response
		res, err = http.Get(url)
		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
	}()

	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

func test3() {
	f := Future(test3_ex)
	fmt.Println("V3 is this locked again")
	result, err := f()
	body := result.([]byte)
	if err == nil {
		fmt.Printf("V3 response length %d\n", len(body))
	} else {
		fmt.Printf("V3 error is %v\n", err)
	}
}

func test3_ex() (interface{}, error) {
	var res *http.Response
	res, err := http.Get("https://api.github.com/users/octocat/orgs")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// Future boilerplate method
func Future(f func() (interface{}, error)) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		result, err = f()
	}()

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
