package connect

import (
	"net/http"
	"fmt"

	"github.com/nickmancari/sype/pkg/color"
)

func Request(arg string) {

	resp, err := http.Get(arg)
	if err != nil {
		fmt.Println(err)
	}

	serverInfo := resp.Header.Get("server")

	if serverInfo == "" {
		fmt.Println(color.Red+"No Server Info Found."+color.Reset)
	} else {
		fmt.Println(color.Green+serverInfo+color.Reset)
	}

	moreServerInfo := resp.Header.Get("X-Backend-Server")
	fmt.Println(moreServerInfo)

}
