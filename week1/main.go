package main

import (
	"fmt"
	"math"
	"slices"
)

func SumInts(arr []int) int {
	sum := 0

	if len(arr) == 0 {
		return sum
	}

	for _, v := range arr {
		sum += v
	}

	return sum
}

func FindMax(arr []int) int {
	max := math.MinInt + 1

	for _, v := range arr {
		max = int(math.Max(float64(max), float64(v)))
	}

	return max
}

func ReverseSlice(arr []int) {
	arrLen := len(arr)

	for i := 0; i < arrLen/2; i++ {
		temp := arr[i]
		arr[i] = arr[arrLen-i-1]
		arr[arrLen-i-1] = temp
	}
}

func RemoveAt(arr []int, removeIndex int) ([]int, bool) {

	if removeIndex < 0 || len(arr) <= removeIndex {
		return arr, false
	}
	reuslt := []int{}

	return append(reuslt, append(arr[:removeIndex], arr[removeIndex+1:]...)...), true
}

func IsAnagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	mp := map[rune]int{}

	unique := 0

	for _, v := range s1 {
		repeat, ok := mp[v]
		if !ok {
			unique++
			mp[v] = 1
			continue
		}
		mp[v] = repeat + 1
	}

	for _, v := range s2 {
		repeat, ok := mp[v]
		if !ok {
			return false
		}
		mp[v] = repeat - 1

		if repeat-1 == 0 {
			unique--
		}
	}

	return unique == 0
}

func RotateSlice(arr []int, k int) { // pick left put right
	if k <= 0 || len(arr) == 0 {
		return
	}
	rotate := k % len(arr)
	if rotate == 0 {
		return
	}

	if rotate == 0 {
		return
	}

	left := arr[:len(arr)-rotate]
	right := arr[len(arr)-rotate:]

	slices.Reverse(left)
	slices.Reverse(right)

	slices.Reverse(arr)

	fmt.Println(arr)
}

func Deduplicate(arr []int) []int {

	result := []int{}

	unique := map[int]bool{}
	for _, v := range arr {
		if _, ok := unique[v]; !ok {
			result = append(result, v)
			unique[v] = true
		}
	}

	return result
}

func GroupByLength(word []string) map[int][]string {

	mp := map[int][]string{}
	for _, v := range word {

		sl, ok := mp[len(v)]
		if !ok {
			mp[len(v)] = []string{v}
			continue
		}

		sl = append(sl, v)
		mp[len(v)] = sl
	}

	return mp
}

func WordFrequency(word string) map[string]int {

	mp := map[string]int{}
	for _, ru := range word {
		runeSting := string(ru)
		repeat, found := mp[runeSting]

		if !found {
			mp[runeSting] = 0
			continue
		}

		mp[runeSting] = repeat + 1
	}

	return mp
}

func MergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		result[k] += v
	}

	return result
}
func main() {
	source := []int{1, 2, 6234, 2345, 346, 56, 7, 7, 4, 3, 7, 63, 34}

	sum := SumInts(source)

	fmt.Println("Result of SumInts - ", sum)

	max := FindMax(source)
	fmt.Println("Result of FindMax - ", max)

	ReverseSlice(source)
	fmt.Println("Result of ReverseSlice - ", source)

	source = []int{1, 2, 6234, 2345, 346, 56, 7, 7, 4, 3, 7, 63, 34}

	result, isRemove := RemoveAt(source, 1)
	fmt.Println("Result of RemoveAt - ", result, isRemove)

	source = []int{1, 2, 7, 7, 1, 2}
	isAnagram := IsAnagram("asgal", "asgla")
	fmt.Println("Result of isAnagram - ", isAnagram)

	source = []int{1, 2, 6234, 2345, 346, 56, 7, 7, 4, 3, 7, 63, 34}
	RotateSlice(source, 5)
	fmt.Println("Result of RotateSlice - ", source) // 3, 7, 63, 34 , 1, 2, 6234, 2345, 346, 56, 7, 7, 4,

	unique := Deduplicate(source)
	fmt.Println("Result of RotateSlice - ", unique)

	word := []string{"a", "bb", "ccc", "dd", "eee", "f", "ggg"}

	ma := GroupByLength(word)
	fmt.Println("Result of RotateSlice - ", ma)

	maps := WordFrequency("ashdgahspuiaepoiashegaksheoiajseahgkjasklgajsdkghalsjhaisuehgaehaikesjhdgoaihsgaf")
	fmt.Println("Result of WordFrequency - ", maps)

	originalMap := map[string]int{"a": 1, "b": 2}
	anotherMap := map[string]int{"b": 3, "c": 4}

	result2 := MergeMaps(originalMap, anotherMap)
	fmt.Println("Result of MergeMaps :", result2) // map[a:1 b:5 c:4]
}
