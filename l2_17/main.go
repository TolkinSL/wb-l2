package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "usage: telnet host port [--timeout=10s]")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	addr := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", addr, *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "connection error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		_, _ = io.Copy(conn, os.Stdin)

		if tcpConn, ok := conn.(*net.TCPConn); ok {
			_ = tcpConn.CloseWrite()
		}
	}()

	_, _ = io.Copy(os.Stdout, conn)

	<-done
}