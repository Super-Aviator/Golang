package main

import "fmt"

const boilingF=212.0

func main() {
	var f=boilingF
	var boilingC=(f-32)*5/9
	fmt.Printf("%g°F is %g °c",f,boilingC)
}
