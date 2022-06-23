package main

import (
	"flag"
	"log"
	"os"
)

var (
	chmod = flag.Bool("chmod", true, "wether to execute chmod")
	chown = flag.Bool("chown", true, "wether to execute chown")
)

func main() {
	flag.Parse()
	file, err := os.Create("/tmp/hello")
	if err != nil {
		log.Fatal(err)
	}
	if *chmod {
		err = os.Chmod(file.Name(), 420)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *chown {
		err = os.Chown(file.Name(), 1000, 1000)
		if err != nil {
			log.Fatal(err)
		}
	}
}
