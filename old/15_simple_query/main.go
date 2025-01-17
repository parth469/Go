package main

import (
	"fmt"
	"net/url"
)

const urls string = "https://api.example.com/v1/users/12345/details?include=posts,comments&sort=created_at:desc&limit=10"

func main() {
	fmt.Println("welcome to simple things")
	result ,err:= url.Parse(urls)
	if err !=nil {
		panic(err)
	} 

	fmt.Println(result.Scheme,result.Host,result.Path,result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams["include"])


}
