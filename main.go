package main

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/rsc/qr"
	"fmt"
	"io/ioutil"
	//	"os"
)

func main() {
	text := uuid.New()
	data, _ := qr.Encode(text, qr.H)
	content := data.PNG()
	fmt.Println(text)
	ioutil.WriteFile(text+".png", content, 0644)
}
