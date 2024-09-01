package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// array := [5]string{"sd", "34", "345"}
	// for as, it := range array {
	// 	fmt.Println(it, as, len(array))
	// }

	// list := []string{"afs", "DSf"}
	// list = append(list, "hi", "sdf1", "sdg")
	// for index, it := range list {
	// 	fmt.Println("sd", it, index)
	// }
	// sort.Slice(list, func(i, j int) bool {
	// 	return len(list[i]) < len(list[j])
	// })
	// i, found := sort.Find(len(list), func(i int) int {
	// 	return strings.Compare("hi", list[i])
	// })

	// if found {
	// 	fmt.Printf("found %s at entry %d\n", "jo", i)
	// } else {
	// 	fmt.Printf("%s not found, would insert at %d", "target", i)
	// }

	// fmt.Println(list, len(list), slices.IsSorted(list))

	course := []string{}
	course = append(course, "reactjs", "nextsj", "nodejs", "express")
	fmt.Println(len(course), course)

	index, found := sort.Find(len(course), func(i int) int {
		fmt.Println(course[i], strings.Compare(course[i], "nodejs"))
		return strings.Compare(course[i], "nodejs")
	})
	fmt.Println(index, found)
	if found {
		course = append(course[:index], course[index+1:]...)
		fmt.Println(len(course), course, "why")
	}

}
