package main

import (
	"bufio"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func BenchmarkTCPServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:18888")
		if err != nil {
			panic(err)
		}
		request, err := http.NewRequest("get", "http://localhost:18888", nil)
		if err != nil {
			panic(err)
		}
		request.Write(conn)
		_, err = http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUDSStreamServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "bench-unixdomainsocket-sample"))
		if err != nil {
			panic(err)
		}
		request, err := http.NewRequest("get", "http://localhost:18888", nil)
		if err != nil {
			panic(err)
		}
		request.Write(conn)
		_, err = http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			panic(err)
		}
	}
}

func TestMain(m *testing.M) {
	go UnixDomainSocketSteamServer()
	go TCPServer()
	time.Sleep(1 * time.Second)

	code := m.Run()
	os.Exit(code)
}
