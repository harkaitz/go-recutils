package main

import (
	"fmt"
	"io"
	"os"

	"github.com/harkaitz/go-recutils"
)

func main() {
	r := recfile.NewReader(os.Stdin)
	n := 0
	for {
		fields, err := r.Next()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		if n > 0 {
			fmt.Println("")
		}
		fmt.Println("Record:", n)
		for _, field := range fields {
			fmt.Printf("%s: %s\n", field.Name, field.Value)
		}
		n++
	}
}
