package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`r([a-z])([a-z])a`)

	fmt.Println(regex.MatchString("rata"))
	fmt.Println(regex.MatchString("rara"))
	fmt.Println(regex.MatchString("raSa"))

	fmt.Println(regex.FindAllString("rata rara rasa rabu rala raga raKa", 10))
}
