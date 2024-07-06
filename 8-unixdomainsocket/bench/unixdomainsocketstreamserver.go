package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UnixDomainSocketSteamServer() {
	path := filepath.Join(os.TempDir(), "bench-unixdomainsocket-sample")
	os.Remove(path)
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	// defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			_, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
