package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		timeNow := time.Now()
		writer.Write([]byte(fmt.Sprintf("%v, %v %v %v", timeNow.Weekday(), timeNow.Day(), timeNow.Month(), timeNow.Year())))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
