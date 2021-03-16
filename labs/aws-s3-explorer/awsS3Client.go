package main

import (
	"bufio"
	"log"
	"fmt"
	"net/http"
	"flag"
)

func main() {

	var P = flag.String("proxy", "localhost:9876", "Proxy")
	var BName = flag.String("bucket", "", "Bucket Name")
	var Dir = flag.String("directory", "", "Directory")
	flag.Parse()
	if *BName == "" {
		log.Fatalln("Incorrect Argument")
	}
	request := fmt.Sprintf("http://%v/example?bucket=%v&dir=%v", *P, *BName, *Dir)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatalln(err)
	}
	
	S := bufio.NewScanner(resp.Body)
	for i := 0; S.Scan(); i++ {
		log.Println(S.Text())
	}
	resp.Body.Close()
}