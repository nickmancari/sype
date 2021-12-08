package connect

import (
	"net/http"
	"fmt"
)

func Request(arg string) {
	resp, err := http.Get(arg)
	if err != nil {
		fmt.Println(err)
	}

	serverInfo := resp.Header.Get("server")

	if serverInfo == "" {
		fmt.Println("No Server Info Found.")
	} else {
		fmt.Println(serverInfo)
	}

}
