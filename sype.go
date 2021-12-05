package main

import (
	"os"
	"fmt"
	"github.com/nickmancari/sype/connect"
)

func main() {
//	os.Args[1]
	if len(os.Args) > 1 {
		url := os.Args[1]
		connect.Request(url)
	} else {
		fmt.Println("Website Not Given")
	}

}
