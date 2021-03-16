package main
import (
	"encoding/xml"
	"fmt"
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
func main() {

}
