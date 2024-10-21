package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"rangga", "dwi", "mahendra", "doe"}
	values := []int{100, 200, 190, 250, 80}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "eko"))
	fmt.Println(slices.Index(names, "eko"))
}