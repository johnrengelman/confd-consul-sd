// Copyright (c) 2014 John Engelman. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	// "github.com/kelseyhightower/confd/backends"
	"github.com/kelseyhightower/confd/log"
)

var (
	debug        bool
	quiet        bool
	verbose      bool
	transport    string
	port         int
	socket       string
	printVersion bool
)

func main() {
	flag.Parse()
	if printVersion {
		fmt.Printf("confd-consul-sd %s\n", Version)
		os.Exit(0)
	}
	log.Notice("Starting confd-consul-sd")
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Error("Could not listen on port 6000")
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleRequest(conn)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
	flag.BoolVar(&quiet, "quiet", false, "enable quiet logging")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.StringVar(&transport, "transport", "tcp", "the transport to listen on (tcp or unix)")
	flag.IntVar(&port, "port", 6000, "the tcp port to listen on")
	flag.StringVar(&socket, "socket", "/var/run/confd-consul-sd.sock", "Unix socket to listen on")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
}

func handleRequest(conn net.Conn) {
	log.Info("Handling request")
	conn.Write([]byte("{\"status\": \"ok\"}"))
	conn.Close()
}
