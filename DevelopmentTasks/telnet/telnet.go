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
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
			os.Exit(1)
		}
	}()

	if _, err := io.Copy(conn, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
		os.Exit(1)
	}
}
