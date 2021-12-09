package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/nickmancari/sype/connect"
)

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		if strings.Contains(url, "https://") == true {
			connect.Request(url)
		} else {
			connect.Request("https://"+url)
		}
	} else {
		fmt.Println("Website Not Given")
	}

}
