package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//Channels and interfaces (And how it works in memory) -> Next talk :)
//Problems for next sessions
/*

1- Check "VIP" -> Http (?)
2-
3-
4-
5-


*/

//Example file server running on port 1234

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + (r.URL.Path)[1:]))
}

func usage() {
	fmt.Println("The usage of this command is =)... ")
}

func main() {

	url := flag.String("url", "http://www.google.es", "This is the URL to check.")
	rcode := flag.Int("code", 200, "This is the expected HTTP response code.")
	content := flag.String("content", "aaa", "This is the expected content.")
	flag.Parse()

	r, e := http.Get(*url)
	if e != nil {
		log.Panic(e)
	}

	if r.StatusCode != *rcode {
		fmt.Printf("Unexpected response status %d, expected %d", r.StatusCode, *rcode)
	} else {
		fmt.Printf("OK")
	}

	b, _ := ioutil.ReadAll(r.Body)
	if strings.Contains(string(b[:]), *content) {
		fmt.Printf("The URL %s contains the string %s", *url, *content)
	}
	return
	//Get the parameters given by the user or whatever
	//thing that's running the code, start at 1 to ignore the file
	argv := os.Args[1:]

	//Print params values
	fmt.Println(argv)

	//Can't have unused variables :(
	//"foreach" statement
	for testIndex, v := range argv {
		fmt.Println(testIndex)
		fmt.Println(v)
	}

	//fmt.Println("I'm main.")
	//http.ListenAndServe(":1235", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":1234", nil)
}
