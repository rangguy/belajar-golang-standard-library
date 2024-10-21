package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rangga Dwi", "Rangga"))
	fmt.Println(strings.Split("Rangga Dwi", " "))
	fmt.Println(strings.ToLower("Rangga Dwi"))
	fmt.Println(strings.ToUpper("Rangga Dwi"))
	fmt.Println(strings.Trim(" Rangga Dwi  ", " "))
	fmt.Println(strings.ReplaceAll("Tri Tri Dwi Mahendra ", "Tri", "Rangga"))

}
