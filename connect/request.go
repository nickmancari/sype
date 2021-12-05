package connect

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Request(arg string) {
	req, _ := http.NewRequest("GET", arg, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
