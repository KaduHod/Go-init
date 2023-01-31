package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
) 

func main(){
	var one number = 1;
	var two number = 2;

	var message string = fmt.Sprintf("%v + %v = %v",one, two, one + two)

	fmt.Println(message)
}

func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}