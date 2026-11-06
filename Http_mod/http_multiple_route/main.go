package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json");
	_,_=w.Write([]byte("Welcome to our Company /me?name=Asish"));

	// _,_=w.Write([]byte(`{"msg":"Welcome to our Company"}`));
}

func meHandler(w http.ResponseWriter,r*http.Request){
	w.Header().Set("Content-Type","application/json");
	name:=r.URL.Query().Get("name");

	if name==""{
		name="Guest"
	}
	_,_=w.Write([]byte(name))
	// _,_=w.Write([]byte(`{"msg":"Hi !! I am founder of Crix Solution"}`));

}


func main(){

	http.HandleFunc("/",rootHandler);
	http.HandleFunc("/me",meHandler);


	err:=http.ListenAndServe(":8901",nil);
	fmt.Println(err);
}
