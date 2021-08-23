package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	claimReader := ClaimReader{filename: filename}
	listOfClaims := claimReader.ReadClaims()
	claims := Claims{content: listOfClaims}
	fmt.Println(claims)
}
