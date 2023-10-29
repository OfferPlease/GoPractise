package main

import "fmt"

// take some number of slices return the sums of all elements except the first in them
// can take out some parts of slices with (slice[low: high])
func SumAllTails(numsSlices ...[]int) []int {
	var sums []int
	for _, nums := range numsSlices {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumSlice(nums[1:]))
		}
	}
	return sums

}

// take some number of slices, return the sums of all elements in them in a slice
func SumAll(numSlices ...[]int) []int {
	// var sums []int
	// for _, numbers := range numSlices {
	// 	sums = append(sums, SumSlice(numbers))
	// }
	var sums = make([]int, len(numSlices))
	for i, numbers := range numSlices {
		sums[i] = SumSlice(numbers)
	}
	return sums
}

func SumSlice(numbers []int) (sum int) {
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}
func Hello() string {
	return "Hello world\n"
}

func main() {
	fmt.Print(Hello())
}
