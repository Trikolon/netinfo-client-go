package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Disable logger timestamp prefix
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Endpoint and path to fetch network info from
	endpoint := "https://net.pbz.im"
	path := "/ip" // Only fetch and show ip by default

	// If the user provides a path, use it for the request
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// Contact endpoint with set path
	response, err := http.Get(endpoint + "/" + path)
	if err != nil {
		log.Fatalln("Error while getting network info from "+endpoint, err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalln("Unexpected http status from endpoint", response.Status)
	}

	// Read body as string
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error while processing endpoint response", err)
	}

	// Print body
	log.Println(string(body))
}
