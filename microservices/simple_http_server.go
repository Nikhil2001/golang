package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type helloWorld struct {
	Message string `json:"message"`
	Count int
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	hello:= helloWorld{"Hello World",6}
	data, _ :=json.Marshal(hello) 
	fmt.Fprint(w,string(data))
}

func main(){
 
	port := 8080

	http.HandleFunc("/helloworld",helloWorldHandler)
	log.Printf("starting Server on %v\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}