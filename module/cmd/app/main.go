package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main(){
	msg:=greet.Hello("Asish");
	fmt.Println(msg);
}