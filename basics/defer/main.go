package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func t1() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func t2() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func t3() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func t4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	t4()
}
