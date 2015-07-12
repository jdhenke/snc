package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please specify a host.\n")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Printf("Too many arguments.\n")
		os.Exit(1)
	}
	host := os.Args[1]
	port := "443"
	if len(os.Args) == 3 {
		port = os.Args[2]
	}
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", host+":"+port, config)
	if err != nil {
		fmt.Printf("Error dialing host: %v\n", err)
		os.Exit(1)
	}
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
