package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	indexp:="/hello/3"
	server:=http.Server{
		Addr:"127.0.0.1:8081",
		//Handler:nil,
	}
	//http.HandlerFunc("/hello",hello)
	//http.HandlerFunc("/hi",hi)
	http.Handle("/",http.HandlerFunc(firstpage))
	http.Handle("/hi",http.HandlerFunc(hi))
	http.Handle("/hello/2",http.HandlerFunc(hello2))
	http.Handle(indexp,http.HandlerFunc(hello3))

	err:=server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
func firstpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"main!")
}
func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello2!")
}
func hello3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello3!")
}
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hi!")
}
