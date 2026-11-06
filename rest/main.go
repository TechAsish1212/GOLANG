package main

import (
	"fmt"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello world")
}

func main(){
	http.HandleFunc("/hello",Helloworld);
	fmt.Println("Server on started at 8080");
	http.ListenAndServe(":8080",nil);
}