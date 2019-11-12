# Go Lang Tutorials

## Data Tutorial 

1. Downloading the Go Compiler and it's libraries.
[Go Download](https://golang.org/dl/)

2. Basic Go template
```Go
package main

import ("fmt")

func main() {
    fmt.Println("GOOO!!!")
}
```
3. import the packages.

```Go
import ( "fmt"
"net/http"
"io/ioutil"
"encoding/xml")
```
4. Implementing
- Get response from washingtonpost sitemap .xml doc
- Passing resp.body to bytes variable
- Adding it to string variable to print to console.
- Closing the response
```Go
func main() {
    
    resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")

    bytes, _ := ioutil.ReadAll(resp.Body)
    
    string_body := string(bytes)
    fmt.Println(string_body) 

    resp.Body.Close()
    
}
```
5. Needed Functions
- Struct SiteMapIndex
- what []type == slice
- For both slices of Location and Location itself.
```Go
type SiteMapIndex struct {
    Locations []Location `xml:"sitemap"`
}

type Location struct {
    Loc string `xml:"loc"`
}
```
