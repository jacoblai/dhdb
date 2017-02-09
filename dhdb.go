package main

import (
	"Redico"
	"log"
	"os"
	"path/filepath"
	"flag"
	"fmt"
	"runtime"
	"os/signal"
	"strconv"
)

var dhdbVersion = "0.0.1"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		port = flag.Int("p", 6380, "dhdb server host port number")
		pass = flag.String("a", "icoolpy.com", "dhdb client auth password")
	)
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	redServer, err := Redico.Run(dir,strconv.Itoa(*port))
	if err != nil {
		panic(err)
	}
	defer redServer.Close()
	redServer.RequireAuth(*pass)
	fmt.Println("DHDB Version:", dhdbVersion)
	fmt.Println("DHDB Port:", strconv.Itoa(*port))
	fmt.Println("Power By ICOOLPY.COM")

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			close(cleanupDone)
		}
	}()
	<-cleanupDone
	fmt.Println("\nStoped DHDB...\n")
}