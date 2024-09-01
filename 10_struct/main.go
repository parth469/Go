package main

import "fmt"

func UserInput() (selectValue int) {
	fmt.Printf("Please Select YOu dice Number You get :")
	fmt.Scan(&selectValue)
	return
}

func GameRule(sum int, userChoise int) (subSum int) {
	switch userChoise {
	case 1:
		subSum += 1
	case 2:
		subSum += 2
	case 3:
		subSum += 3
	case 4:
		subSum += 4
	case 5:
		subSum += 5
	case 6:
		subSum += 6
	default:
		fmt.Println("Please select proper value")
		break
	}
	return
}
func main() {
	sum, term := 0, 0
	for sum < 100 {
		userValue := UserInput()
		sum += GameRule(sum, userValue)
		fmt.Println("you current sum is", sum)
		term++
	}
	fmt.Println("You win this game with", term, "Trun", sum)

}
