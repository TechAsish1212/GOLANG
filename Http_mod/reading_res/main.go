package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	url:="https://dummyjson.com/todos";
	resp,err:=http.Get(url);
	if err!=nil{
		fmt.Println(err);
		return;
	}

	defer resp.Body.Close();

	fmt.Println("Status Code:: ",resp.StatusCode);
	fmt.Println("status:: ",resp.Status);

	bodyBytes,err:=io.ReadAll(resp.Body);

	if err!=nil{
		fmt.Println(err);
		return;
	}

	bodyText:=string(bodyBytes);
	max:=250;
	if(len(bodyText)<max){
		max=len(bodyText);
	}

	fmt.Println("Response::: -> ",bodyText[:max])
}