package main

import ( "fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SiteMapIndex struct {

	// slice []type == slice
	Locations []Location `xml:"sitemap"`

} // end SiteMapIndex struct

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	// get response from washingtonpost sitemap .xml doc
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")

	// passing resp.body to bytes variable
	bytes, _ := ioutil.ReadAll(resp.Body)
	/* adding it to string variable to print to console.
	string_body := string(bytes)
	fmt.Println(string_body) 
	*/

	// closing response to end session
	resp.Body.Close()
	
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

} // end main function