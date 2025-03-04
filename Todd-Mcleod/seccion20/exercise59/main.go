package main

import "fmt"


func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(foo(xi...))

	fmt.Println(bar([]int{1,2,3}))

}



func foo(nums ...int) int {
	total := 0
	for _, v := range nums{
		total = total+v
	}

	return total
}

func bar(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}


