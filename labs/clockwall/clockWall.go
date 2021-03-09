package main

import (
	"io"
	"net"
	"os"
	"strings"
	"strconv"
)

func MakeRequest(str string, ch chan int) {
	conn, err := net.Dial("tcp", str)
	if err == nil {
		newStrInt:=strings.Split(str, ":")
		num, err := strconv.Atoi(newStrInt[1])
		if err == nil {
			defer conn.Close()
			_, err = io.Copy(os.Stdout, conn)
			if err == nil {
				ch <- num
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	var connections []string

	for _,i := range args {
		newDiv:=strings.Split(i, "=")
		newConnection := newDiv[1]
		connections = append(connections, newConnection)
	}

	// Buffered chanel with all the conections
	ch := make(chan int, len(connections)) 
	for _,conn := range connections {
		go MakeRequest(conn, ch)
	}
	<- ch
	close(ch)
}