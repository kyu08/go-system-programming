package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, _ := os.Create("test.csv")
	w := csv.NewWriter(file)
	w.Write([]string{"name", "age"})
	w.Write([]string{"taro", "3"})
	w.Flush()
}
