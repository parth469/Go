package main

import "fmt"

type Size string

const (
	small      Size = "small_Size"
	medium     Size = "medium"
	large      Size = "large"
	extraLarge Size = "extraLarge"
)

func main() {
	fmt.Println(small, medium, large, extraLarge)
	sli := make([]int, 0)
	var slis []int
	sliscs := new([]int)
	fmt.Println(sli == nil)
	fmt.Println(slis == nil)
	fmt.Println(sliscs == nil)

	mp := make(map[string]int)
	mp["ad"] = 23
	mp["adb"] = 23
	mp["ads"] = 233

	ok, isExist := mp["adg"]

	fmt.Println(ok, isExist)

	fmt.Println(mp)

	sli = append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(sli)
}
