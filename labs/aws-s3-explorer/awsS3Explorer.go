package main
import (
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)
//General Data Structure
type GenData struct {
	BName string
	Objs map[string]bool
	Dirs map[string]bool
	Exts map[string]int
}
//Specific Data Structure
type GenData struct {
	SName string
	Objs map[string]bool
	Dirs map[string]bool
	Exts map[string]int
}
//Content
type Content struct {
	Key string `xml:"Key"`
}

//ListBucket
type ListBucket struct {
	XMLName  xml.Name  `xml:"ListBucketResult"`
	Contents []Content `xml:"Contents"`
}
func search(w http.ResponseWriter, r *http.Request)
{
	max:=5
	exts := make(map[string]int)
	dirs := make(map[string]bool)
	objs := make(map[string]bool)
	
	resp, err := http.Get(fmt.Sprintf("https://%v.s3.amazonaws.com", r.FormValue(BName)))
	if err!=""{
		log.Fatalln(err)
	}
}
func main() {
	var p = flag.Int("port", 9876, "Port")
	flag.Parse()
	s := fmt.Sprintf("localhost:%v", *p)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		search(w, r)
    })
	
	err := http.ListenAndServe(socket, nil)
	if err != nil {
		log.Fatal(err)
	}
}
