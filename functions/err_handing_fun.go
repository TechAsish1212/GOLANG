package main

import "fmt"

func div(a,b float32 ) (float32,error){
	if(b==0){
		return 0,fmt.Errorf("denomintor must be zero");
	}
	return a/b,nil;
}

func main(){
	fmt.Println("Error Handle");
	res,_:=div(17,0);
	fmt.Println("ans:: ",res);
}
