package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// struktur data ring, circular list

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++{
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// untuk iterasi setiap datanya
	data.Do(func (value any)  {
		fmt.Println(value)
	})
}