package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// reader
	input := strings.NewReader("ini kalimat panjang dengan dua baris\nini baris kedua")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// writer
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world!\n")
	_, _ = writer.WriteString("selamat belajar\n")

	writer.Flush()

}
