package main

import (
	"flag"
	"log"
)

func main() {
	dc := &DoitConfig{}

	port := flag.String("p", "8080", "Port")
	config := flag.String("c", "", "Load config file")
	serverMode := flag.Bool("s", false, "Enable server mode")

	flag.Parse()

	log.Println(*port, *serverMode, *config)

	if *config != "" {
		//load config
		err := dc.Read(*config)
		if err != nil {
			log.Fatal("Unable to load config file", err)
		}
	} else {
		//manual load config
		log.Fatal("Unable to load config")
	}

	if *serverMode {
		storage, err := NewStorage(dc.Storage.Type, dc.Storage.Location)
		ds := &DoitServer{Store: storage}
		err = ds.Listen(port, dc)
		if err != nil {
			log.Fatal("Unable to listen on specified port:", err)
		}
	}
}
