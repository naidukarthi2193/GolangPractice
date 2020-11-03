package main

import (
	"fmt"
	"sort"
)

//Person ...
type Person struct {
	name string
	age  int
}

//ByAge ...
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	var nums []int
	nums = append(nums, 1, 2, 3, 4, 8, 4, 7, 67, 2, 4, 675, 244, 6, 4534, 5645, 7, 4)
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	peeps := []Person{{"Karthik", 23}, {"Thiu", 1023}, {"shubhum", 20}, {"venky", 24}}
	sort.Sort(ByAge(peeps))
	fmt.Println(peeps)
}
