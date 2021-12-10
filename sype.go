package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/nickmancari/sype/connect"
	"github.com/nickmancari/sype/pkg/color"
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
		fmt.Println(color.Yellow+"Website Not Given"+color.Reset)
	}

}
