package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)

func Test() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(index)))
	http.Handle("/hello", timeMiddleware(http.HandlerFunc(hello)))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//中间件 业务剥离
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		//logger.Println(next,timeElapsed)
		logger.Println(r.RequestURI,timeElapsed)
	})
}

func hello(wr http.ResponseWriter, r *http.Request) {
	for j := 0; j < 5000; j++ {
		wr.Write([]byte("hello"))
	}
}

func index(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("index"))
}
