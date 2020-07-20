package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func calculator(st []string, c net.Conn) {
	xdecimal, _ := decimal.NewFromString(st[0])
	ydecimal, _ := decimal.NewFromString(st[2])
	if st[1] == "+" {
		fmt.Fprintln(c, "\t", xdecimal.Add(ydecimal))
	}
	if st[1] == "-" {
		fmt.Fprintln(c, "\t", xdecimal.Sub(ydecimal))
	}
	if st[1] == "*" {
		fmt.Fprintln(c, "\t", xdecimal.Mul(ydecimal))
	}
	if st[1] == "/" {
		fmt.Fprintln(c, "\t", xdecimal.Div(ydecimal))
	}
}
func echo(c net.Conn, shout string, delay time.Duration) {
	num := strings.Split(shout, " ")
	calculator(num, c)
	//	fmt.Fprintln(c, "\t", res)
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
