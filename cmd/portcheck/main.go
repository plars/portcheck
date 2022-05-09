/*
Simple tool to check whether a port is active, and keep checking
it like ping

Written by:
  Paul Larson <pwlars@gmail.com>
*/

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

type args struct {
	host  string
	port  string
	count int
}

func getArgs() args {
	var args args
	flag.Usage = usage
	count := flag.Int("count", 1, "Number of packets to send")
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
		os.Exit(1)
	}
	args.host = flag.Arg(0)
	args.port = flag.Arg(1)
	args.count = *count
	return args
}

func usage() {
	fmt.Printf("Usage: %s [options] HOST PORT\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	args := getArgs()
	for i := 0; ; {
		i++
		fmt.Printf("%d %s:%s - %s\n", i, args.host, args.port, connect(args.host, args.port))
		if i >= args.count {
			break
		}
		time.Sleep(time.Second)
	}
}

func connect(host, port string) string {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return "Down"
	}
	if conn == nil {
		return "Down"
	}
	conn.Close()
	return "Up"
}
