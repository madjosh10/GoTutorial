package main

//import "fmt" and "net/http"
import ("fmt"
		"net/http")

//create the indexHandler function to write and read the http request
func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,`
			<h1>hey there</h1>
			<p>Go is fast! </p>
			<p>Learning how to set up a local server with Go</p>
		`)
	fmt.Fprintf(w,"<p>You %s even add %s </p>", "can", "<strong>variables</strong>")

}

// make the main method to handle the appropriate function for its path.
// ListenAndServe is the port # for the local host on your machine
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}