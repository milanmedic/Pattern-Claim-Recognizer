package main

import (
	"fmt"
	"os"
	"strings"
)

func HandleError(err error, action string) {
	fmt.Println(err.Error())
	if strings.EqualFold(action, "terminate") {
		os.Exit(1)
	}
}

func ByteArrToStr(input []byte) string {
	return string(input)
}
