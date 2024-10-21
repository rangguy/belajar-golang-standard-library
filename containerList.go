package main

// struktur data doubly linked list

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rangga")
	data.PushBack("Dwi")
	data.PushBack("Mahendra")

	head := data.Front()
	fmt.Println(head.Value) // rangga

	next := head.Next()
	fmt.Println(next.Value) // dwi

	next = next.Next()
	fmt.Println(next.Value) // mahendra

	for e := data.Front(); e != nil; e = e.Next() {
		// e -> head; selama masih ada elemen; diambil elemen selanjutnya
		fmt.Println(e.Value)
	}
}