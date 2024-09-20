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

	cr := time.Date(1034, time.April, 2524572345723894759, 2, 3, 5, 0, time.Local)
	fmt.Println(cr)

	start := time.Now()
	a := 1
	for i := 0; i < 19994745; i++ {
		a++
	}
	el := time.Since(start)
	fmt.Println(el)
}
