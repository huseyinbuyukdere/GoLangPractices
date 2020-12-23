package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Status Code:", resp.StatusCode)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
