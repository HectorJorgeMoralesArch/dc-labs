// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"flag"
	"os"
)
//Returns the time at the 'name' location
func GPS(t time.Time, name string) (time.Time) {
	loc, err := time.LoadLocation(name)
	if err!=nil{
		panic(err)
	}
	return t.In(loc)
}
//Generates the Output
func Output(ch chan string, place string){
	for {
		time.Sleep(1 * time.Second)
		ch <- GPS(time.Now(), place).Format("15:04:05\n")
	}
}

func handleConn(c net.Conn, place string) {
	defer c.Close()
	ch:=make(chan string)
	go Output(ch, place)
	for i := range ch {
		str := fmt.Sprintf("%v\t:\t%v", place, i)
		_, err := io.WriteString(c, str)
		if err != nil {
			return 
		}
	}
}

func main() {
	tz := os.Getenv("TZ")
	var port = flag.Int("port", 8010, "Port number.")
	flag.Parse()

	lh := fmt.Sprintf("localhost:%v", *port)
	listener, err := net.Listen("tcp", lh)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("======== Started clock server [%v] at port [%v] ========\n", tz, *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn, tz) // handle connections concurrently
	}
}
