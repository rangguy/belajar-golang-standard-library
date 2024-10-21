package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"rangga", "dwi", "mahendra"})
	_ = writer.Write([]string{"dwi", "kurnia", "wana"})
	_ = writer.Write([]string{"joko", "wi", "dodo"})

	writer.Flush()
}