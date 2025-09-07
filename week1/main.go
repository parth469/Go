package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
	"time"
)

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) PopFront() (T, bool) {
	var ZeroValue T
	if len(q.queue) == 0 {
		return ZeroValue, false
	}

	ZeroValue = q.queue[0]
	q.queue = slices.Delete(q.queue, 0, 1)
	return ZeroValue, true
}

func (q *Queue[T]) PopLast() (T, bool) {
	var ZeroValue T
	leng := len(q.queue)
	if leng == 0 {
		return ZeroValue, false
	}

	ZeroValue = q.queue[leng-1]
	q.queue = q.queue[:leng-1]
	return ZeroValue, true
}

func (q *Queue[T]) PushLast(value T) {
	q.queue = append(q.queue, value)
}

func (q *Queue[T]) PushFront(value T) {
	q.queue = slices.Insert(q.queue, 0, value)
}

func (q *Queue[T]) Front() (T, bool) {
	var zeroValue T
	if len(q.queue) == 0 {
		return zeroValue, false
	}
	return q.queue[0], true
}

func (q *Queue[T]) Back() (T, bool) {
	var zeroValue T
	if len(q.queue) == 0 {
		return zeroValue, false
	}
	return q.queue[len(q.queue)-1], true
}
func QueueStart() {
	var q Queue[int]
	q.PushLast(1)
	q.PushFront(2)
	q.PushLast(3)

	fmt.Println(q)

	fmt.Println(q.PopFront())
	fmt.Println(q.Front())
	fmt.Println(q.Back())
}

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

type User struct {
	ID   int
	Name string
	Age  int
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

func FilterAdults(UserList []User) {
	filterUserList := []User{}

	for _, v := range UserList {
		if v.Age > 18 {
			filterUserList = append(filterUserList, v)
		}
	}

	sort.Slice(UserList, func(i, j int) bool {
		return UserList[j].Age > UserList[i].Age
	})

	fmt.Println(UserList)
	slices.SortFunc(filterUserList, func(a User, b User) int {
		return a.Age - b.Age
	})

	fmt.Println(filterUserList)
}
func createUser() {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank"}
	var users []User

	for i := 1; i <= 5; i++ {
		user := User{
			ID:   i,
			Name: names[rand.Intn(len(names))],
			Age:  rand.Intn(50),
		}
		users = append(users, user)
	}

	FilterAdults(users)
}

type Product struct {
	Name  string
	Price float64
	Tags  []string
}

func FindByTag(ProductList *[]Product, FoundTag string) []Product {
	ProductListValue := *ProductList
	Findproduct := make([]Product, 0, len(ProductListValue))

	for _, product := range ProductListValue {
		if found := slices.Contains(product.Tags, FoundTag); found {
			Findproduct = append(Findproduct, product)
		}
	}

	return Findproduct
}
func generateProduct() {

	names := []string{
		"Laptop", "Phone", "Tablet", "Headphones",
		"Camera", "Keyboard", "Monitor", "Smartwatch",
	}

	tagPool := []string{
		"Electronics", "Portable", "Gaming", "Office",
		"Wireless", "Premium", "Budget", "New",
	}

	var products []Product

	for i := 0; i < 5; i++ {
		// random number of tags (1–3)
		numTags := rand.Intn(3) + 1
		tags := make([]string, 0, numTags)
		for j := 0; j < numTags; j++ {
			tags = append(tags, tagPool[rand.Intn(len(tagPool))])
		}

		product := Product{
			Name:  names[rand.Intn(len(names))],
			Price: float64(rand.Intn(900)+100) + rand.Float64(), // price between 100–999
			Tags:  tags,
		}
		products = append(products, product)
	}
}

type Stack[T any] struct {
	ptr []T
	len int
}

func (s *Stack[T]) Push(value T) {
	s.ptr = append(s.ptr, value)
	s.len++
}

func (s *Stack[T]) Pop() {

	if s.len == 0 {
		fmt.Println("Nothing to pop")
		return
	}
	s.ptr = s.ptr[:s.len-1]
	s.len--
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

	createUser()
	generateProduct()
	QueueStart()
	print := func() string {
		fmt.Println("CALLBACK")
		return "HELLO"
	}
	returnVale := CallbackF(print)

	fmt.Println(returnVale)
}

func CallbackF(fn func() string) string {
	time.Sleep(1 * time.Second)
	return fn()
}

type CSlices[T any] struct {
	ptr []T
}

func (sl *CSlices[T]) Append(value T) {
	if cap(sl.ptr) == len(sl.ptr) {
		newSl := make([]T, len(sl.ptr), cap(sl.ptr)*2)
		copy(newSl, sl.ptr)
		sl.ptr = newSl
	}

	sl.ptr = append(sl.ptr, value)

}

func Swip(a, b *int) {
	*a, *b = *b, *a
}

func MaxMin(sl []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for _, v := range sl {

		if min > v {
			min = v
		}

		if max < v {
			max = v
		}
	}
	return min, max
}

type Student struct {
	Name string
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s Student) GetName() string {
	return s.Name
}

func Sum(number ...int) int {
	sum := 0
	for _, v := range number {
		sum += v
	}

	return sum
}
