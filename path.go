package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	// path
	fmt.Println("Path: ")
	fmt.Println(path.Dir("hello/world.go")) // ngambil foldernya -> hello
	fmt.Println(path.Base("hello/world.go")) // nama filenya -> world.go
	fmt.Println(path.Ext("hello/world.go")) // nama extension
	fmt.Println(path.Join("hello", "world", "main.go")) 

	// filepath
	fmt.Println("\nFile Path: ")
	fmt.Println(filepath.Dir("hello/world.go")) 
	fmt.Println(filepath.Base("hello/world.go")) 
	fmt.Println(filepath.Ext("hello/world.go")) 
	fmt.Println(filepath.IsAbs("hello/world.go")) 
	fmt.Println(filepath.IsLocal("hello/world.go")) 
	fmt.Println(filepath.Join("hello", "world", "main.go")) 

}
