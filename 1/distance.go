package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	l1, l2, err := readInput(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(findDistance(l1, l2))
	fmt.Println(findSimilarity(l1, l2))

}

func readInput(input string) ([]int, []int, error) {
	l1 := []int{}
	l2 := []int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, err
		}
		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}
	return l1, l2, nil
}

func findDistance(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)
	distances := []int{}
	for i := range list1 {
		diff := list1[i] - list2[i]
		absDiff := int(math.Abs(float64(diff)))
		distances = append(distances, absDiff)
	}
	var total int
	for _, v := range distances {
		total += v
	}
	return total
}

func howOften(nums []int, toCheck int) int {
	amount := 0
	for _, num := range nums {
		if num == toCheck{
			amount += 1
		}
	}
	return amount
}

func findSimilarity(l1, l2 []int) int {
	total := 0 
	for _, num1 := range l1 {
			total += num1 * howOften(l2, num1)
	}
	return total
}
