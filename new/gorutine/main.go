package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func excuter(id int) {
	fmt.Println(id)
}

func main() {
	for i := 0; i < 10; i++ {
		go excuter(i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumCPU(), "hel")

	fmt.Println("sd" == "sd")
	strings.Contains("sd", "s")

	fmt.Println(strconv.Atoi("123"))
}
