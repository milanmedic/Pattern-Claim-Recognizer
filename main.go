package main

import "fmt"

func main() {
	claimReader := ClaimReader{path: "./", filename: "claims.txt"}
	listOfClaims := claimReader.ReadClaims()
	claims := Claims{content: listOfClaims}
	fmt.Println(claims)
}
