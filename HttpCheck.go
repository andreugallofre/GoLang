package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + (r.URL.Path)[1:]))
}

func usage() {
	fmt.Println("The usage of this command is =)... ")
}

func main() {
	var urlParsed bytes.Buffer
	url := flag.String("url", "http://www.google.es", "This is the URL to check.")
	rcode := flag.Int("code", 200, "This is the expected HTTP response code.")
	content := flag.String("content", "aaa", "This is the expected content.")
	flag.Parse()

	test := *url

	fmt.Println(!strings.Contains(test[:7], "https://"))
	fmt.Println(!strings.Contains(test[:6], "http://"))

	if !strings.Contains(test[0:6], "http://") && !strings.Contains(test[0:7], "https://") {
		urlParsed.WriteString("http://")
		urlParsed.WriteString(*url)
	} else {
		urlParsed.WriteString(*url)
	}

	fmt.Println(urlParsed.String())

	r, e := http.Get(urlParsed.String())
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

}
