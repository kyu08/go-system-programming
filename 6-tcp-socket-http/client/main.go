package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// TCPソケットを初期化
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	// レスポンスを標準出力
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
}
