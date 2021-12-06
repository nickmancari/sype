package connect

import (
	"net/http"
//	"io/ioutil"
	"fmt"
)

func Request(arg string) {
	resp, err := http.Get(arg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Header.Get("server"))

}
