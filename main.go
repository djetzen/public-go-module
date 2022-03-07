package main

import (
	"fmt"
	"public-go-module/hello"
)

func main() {
	fmt.Println(hello.SayHelloFromPrivateModule("Dominik"))
	fmt.Println(hello.SayHelloFromPublicModule("Dominik"))
}

