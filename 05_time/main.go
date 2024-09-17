package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()
	fmt.Println(currentDate.Format("01/02/2006 15/04:05 Mon January"))

	createTiome := time.Date(2020, time.April, 10, 20, 33, 3, 9, time.Local)
	fmt.Printf(createTiome.String())

	dur := time.Duration(10000000000)
	fmt.Println(dur.Minutes(), "hello")

	str, err := time.ParseDuration("15h10m")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str.Hours())
}
