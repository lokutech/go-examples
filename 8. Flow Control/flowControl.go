package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle") // get call last
	fmt.Println("end")

	defer fmt.Println("1. Start") // Defer gets called in last in first out order
	defer fmt.Println("2. Start")
	defer fmt.Println("3. Start")

	webCall()
	// panicEx()
	// panicEx2()
	// panicEx3()
	panicWithDefer()
}

func webCall() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //Good practice to defer the closing
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicEx() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

func panicEx2() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func panicEx3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}

func panicWithDefer() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}
