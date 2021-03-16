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
	searchedData := Data{Bname: r.FormValue(BName), Dirs: make(map[string]bool), Objs: make(map[string]bool), Exts: make(map[string]int)}

	/* Get XML from URL */
	resp, err := http.Get("https://" + searchedData.BName + ".s3.amazonaws.com")
	if err != nil {
		log.Panicln("Get: " + err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Read body: " + err.Error())
	}
	/* Process XML Response into an object */
	var ListBucket ListBucket
	xml.Unmarshal(data, &ListBucket)
	/* Process Object Data */
	for _, Content := range ListBucket.Contents {
		key := Content.Key
		if strings.Contains(key, ".") {
			if _, exists := searchedData.Objects[key]; !exists {
				searchedData.Objects[key] = true
			}
			container := strings.Split(key, ".")
			ext := container[len(container)-1]
			_, exists := searchedData.Extentions[ext]
			if !exists {
				searchedData.Extentions[ext] = 0
				searchedData.Extentions[ext]++
			}
			else {
				searchedData.Extentions[ext]++
			}
		}
		if strings.HasSuffix(key, "/") {
			if !searchedData.Dirs[key] {
				searchedData.Dirs[key] = true
			}
		}
	}
	print(searchedData)
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
