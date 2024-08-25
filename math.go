package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.30)) // dibulatkan ke atas
	fmt.Println(math.Floor(1.30)) // dibulatkan ke bawah
	fmt.Println(math.Round(1.30)) // dibulatkan ke yang paling dekat
	fmt.Println(math.Max(1.30, 1.29))
	fmt.Println(math.Min(1.30, 1.29))
}
