package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	testData := []byte("0123456789ABCDEF")
	testPath := filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0o644)
	if err != nil {
		panic(err)
	}

	// メモリにマッピング
	f, err := os.OpenFile(testPath, os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	// mmapしたデータに変更を加えてfileにflush
	m[9] = 'X'
	m.Flush()

	// ファイルに変更が加わっていることを確認
	fileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
