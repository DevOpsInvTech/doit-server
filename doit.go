package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime/pprof"

	log "github.com/Sirupsen/logrus"
)

func main() {

	dc := &DoitConfig{}

	port := flag.String("p", "8080", "Port")
	config := flag.String("c", "", "Load config file")
	serverMode := flag.Bool("s", false, "Enable server mode")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write memory profile to this file")

	var f *os.File

	flag.Parse()
	if *cpuprofile != "" {
		ff, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(ff)
	}
	if *memprofile != "" {
		var err error
		f, err = os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)

	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			pprof.StopCPUProfile()
			f.Close()
			os.Exit(0)
		}
	}()

	log.Println(*port, *serverMode, *config)

	if *config != "" {
		//load config
		err := dc.Read(*config)
		if err != nil {
			log.Fatalln("Unable to load config file", err)
		}
	} else {
		//manual load config
		log.Fatal("Unable to load config")
	}

	if *serverMode {
		storage, err := NewStorage(dc.Storage.Type, dc.Storage.Location)
		if err != nil {
			log.Fatalln("Unable to connect to storage:", err)
		}
		ds := &DoitServer{Store: storage}
		err = ds.Listen(port, dc)
		if err != nil {
			log.Fatalln("Unable to listen on specified port:", err)
		}
	}
}
