package main
import (
	"encoding/xml"
	"fmt"
	"flag"
	"log"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)
//General Data Structure
type Data struct {
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
func search(w http.ResponseWriter, r *http.Request){
	searchedData := Data{BName: r.FormValue(BName), Dirs: make(map[string]bool), Objs: make(map[string]bool), Exts: make(map[string]int)}
	dir:=r.FormValue(Dir)
	hasDir:=false
	if dir!=""{
		hasDir=true
	}
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
		if hasDir == true && !strings.HasPrefix(key, dir){
			continue
		}
		if strings.Contains(key, ".") {
			if _, exists := searchedData.Objs[key]; !exists {
				searchedData.Objs[key] = true
			}
			container := strings.Split(key, ".")
			ext := container[len(container)-1]
			_, exists := searchedData.Exts[ext]
			if !exists {
				searchedData.Exts[ext] = 0
				searchedData.Exts[ext]++
			}else {
				searchedData.Exts[ext]++
			}
		}
		if strings.HasSuffix(key, "/") {
			if !searchedData.Dirs[key] {
				searchedData.Dirs[key] = true
			}
		}
	}
	print(searchedData,hasDir)
}
func print(data Data, dir bool) {
	if dir{
		fmt.Println("AWS S3 Explorer\nBucket Name\t\t: " + data.BName)
	}else{

		fmt.Println("AWS S3 Explorer\nDirectory Name\t\t: " + data.BName)
	}
	fmt.Println("Number of objects\t: " + strconv.Itoa(len(data.Objs))+"\nNumber of directories\t: " + strconv.Itoa(len(data.Dirs))+"\nExtensions\t\t: ")
	for key, value := range data.Exts {
		fmt.Print(key + "(" + strconv.Itoa(value) + ")\n")
	}
}
func main() {
	var p = flag.Int("port", 9876, "Port")
	flag.Parse()
	s := fmt.Sprintf("localhost:%v", *p)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		search(w, r)
    })
	
	err := http.ListenAndServe(s, nil)
	if err != nil {
		log.Fatal(err)
	}
}
