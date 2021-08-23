package main

import (
	"regexp"

	"github.com/google/uuid"
)

type ClaimReader struct {
	path     string
	filename string
}

func (cr ClaimReader) ReadClaims() []Claim {
	path := cr.path + cr.filename
	var file File = File{name: path}
	file.OpenFile()
	allContent := file.ReadFile()
	claims := ExtractClaims(allContent)
	return claims
}

func ExtractClaims(content string) []Claim {
	regex := regexp.MustCompile("/\n\\d{1,}\\./")
	unformattedClaims := regex.Split(content, -1)
	allClaims := FormatClaims(unformattedClaims)
	return allClaims
}

func FormatClaims(unformattedClaims []string) []Claim {
	var claims []Claim
	var id string = uuid.New().String()
	for _, unformattedClaim := range unformattedClaims {
		var claim Claim = Claim{id: id, content: unformattedClaim}
		claims = append(claims, claim)
	}
	return claims
}
