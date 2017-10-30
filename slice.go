package gosystem

import (
	"fmt"
	"math/rand"
)

func SliceArray() {
	a := [5]int{1, 2, 3, 4, 5}
	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	myslice = append(myslice, 9, 10, 11)
	fmt.Println(myslice)

	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5
	fmt.Println(s)
	s1 := s[1:2]
	fmt.Println(s1[:cap(s1)])
	s2 := make([]int, 100)
	assignarray(s2)
	s3 := make([]int, 10)
	s4 := make([]int, 10)

	s3 = getstwo(s2, 0, 9)
	fmt.Println(s3)

	fmt.Println(s3[:cap(s3)])

	s4 = getstwocorrect(s2, 0, 9)

	fmt.Println(s4)

	fmt.Println(s4[:cap(s4)])

}

func getstwo(s []int, fi, si int) []int {
	return s[fi:si]
}

func getstwocorrect(s []int, fi, si int) []int {
	output := make([]int, si-fi)
	copy(output, s[fi:si])
	return output
}

func assignarray(input []int) {

	for idx, value := range input {
		input[idx] = rand.Intn(100)
		value = value * 2
	}
}
