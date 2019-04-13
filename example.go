package main

import (
	"github.com/ZacJoffe/clipboard/xclip"
	//"./xclip"
	"fmt"
	"log"
)

func main() {
	err := xclip.Write("test")
	if err != nil {
		log.Fatal(err)
	}

	out, err := xclip.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
