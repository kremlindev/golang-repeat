package main

import (
	"fmt"
)
func main() {

fmt.Println("Write something and I will repeat what you wrote 1000 times.")
var cevap string
fmt.Scanln(&cevap)
for i := 0;i < 1000; i++{
fmt.Println(cevap)
	}
	}
