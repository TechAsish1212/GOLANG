package main

import "fmt"

func evenodd(){
    var num int;
    fmt.Println("Enter the number:: ");
    fmt.Scan(&num);
    if(num%2==0){
        fmt.Println("Given number is a even number");
    }else{
        fmt.Println("Given number is a odd number");
    }
}

func main() {
    fmt.Println("Hello World!")
    evenodd();
}



