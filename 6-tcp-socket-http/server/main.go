package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// TCPソケットを使ったHTTPサーバ
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server is running at localhost:8888\n\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				// タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				// リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					// それ以外はエラーにする
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				// リクエストをデバッグ表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))

				content := "Hello World\n"

				// レスポンスを返す
				response := http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 1,
					// GoのResponse.Write()はHTTP/1.1より古いバージョンが使われる場合、もしくは
					// 長さがわからない場合は`Connection: close`ヘッダーを付与してしまう。
					// これは、クライアントが複数のレスポンスを扱うことから明確にそれぞれのレスポンスを区切る必要があるため。
					// そのためKeep-Aliveを利用するためにはContentLengthを設定する必要がある。
					ContentLength: int64(len(content)),
					Body:          io.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
