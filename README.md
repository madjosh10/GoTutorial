# Go Lang Tutorials

## Data Tutorial 

1. Downloading the Go Compiler and it's libraries.
[Go Download](https://golang.org/dl/)
- Compiling With oppropriate OS
```Windows
go run <filename>.go
```

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
5. Needed Structs
- Struct SiteMapIndex
- Struct Location
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
6. Needed Function
- A function that takes in a Location type and returns a string.
- the String() is the function 'name'
```Go
func (l Location) String() string {
    return fmt.Sprintf(l.Loc)
}
```
7. Last touch ups
- Declaring the s varibale that is of type SiteMapIndex
- Using xml import to Unmarshal the bytes and s. 
- &s is saying point at the address of s.
- printing the s.Locations to the console
```Go
func main() {
    resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
   
    bytes, _ := ioutil.ReadAll(resp.Body)
    
    resp.Body.Close()
    
    var s SiteMapIndex
    xml.Unmarshal(bytes, &s)
    fmt.Println(s.Locations)
    
}
```


