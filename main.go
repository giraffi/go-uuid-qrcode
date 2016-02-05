package main

import (
	"github.com/pborman/uuid"
	qrcode "github.com/skip2/go-qrcode"
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
	data, _ := qrcode.New(text, qrcode.High)
	content := data.PNG()
	fmt.Println(text)
	ioutil.WriteFile(text+".png", content, 0644)
}
