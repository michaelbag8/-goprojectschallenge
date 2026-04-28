package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 3{
		fmt.Println("Usage: code encode/decode")
		return
	}
	
	encodeDecode := os.Args[1]
	code := os.Args[2]

	encodeDecode = strings.ToLower(encodeDecode)

	switch encodeDecode{
	case "encode":
		fmt.Println(encode(code, letterToMorse))
	case "decode":
		fmt.Println(decode(code, letterToMorse))
	default:
		fmt.Println("unknown morse command ")
	}

}