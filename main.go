package main

import (
	"fmt"
	"wxcloudrun-golang/db"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	//
	//log.Fatal(http.ListenAndServe(":80", nil))
}
