package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	service := flag.String("s", "gist", "service name, e.g. gist or imgur")
	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" {
		log.Fatal("undefined filename")
	}

	fmt.Printf("Uploading file %s to %s...\n", filename, *service)

	// Create our client.
	c, err := client(*service)
	if err != nil {
		log.Fatal(err)
	}

	// Read the stream of bytes of the file.
	stream, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Send the stream with the client.
	err = upload(c, stream)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
