package main

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/rsc/qr"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	v := flag.Bool("v", false, "show version")
	flag.Parse()

	if *v {
		fmt.Println("version:0.1.1.dev")
		os.Exit(0)
	}

	text := uuid.New()
	data, _ := qr.Encode(text, qr.H)
	content := data.PNG()
	fmt.Println(text)
	ioutil.WriteFile(text+".png", content, 0644)
}
