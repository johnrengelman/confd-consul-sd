// Copyright (c) 2014 John Engelman. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	// "github.com/kelseyhightower/confd/backends"
	"github.com/johnrengelman/confd-consul-sd/consul"
	"github.com/kelseyhightower/confd/log"
)

var (
	client *consul.Client
)

func main() {
	flag.Parse()
	if printVersion {
		fmt.Printf("confd-consul-sd %s\n", Version)
		os.Exit(0)
	}
	config = Config{
		Transport: "tcp",
		Port:      6000,
		Socket:    "/var/run/confd-consul-sd.sock",
	}
	processFlags()
	// Configure logging.
	log.SetQuiet(config.Quiet)
	log.SetVerbose(config.Verbose)
	log.SetDebug(config.Debug)

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
		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	log.Info("Handling request")
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	prefix, err := reader.ReadString('\n')
	if err != nil {
		log.Error("Error handling request")
		return
	}
	log.Info("Retrieving keys for prefix: " + prefix)

	keys := RetrieveKeys(prefix)
	writer.WriteString(keys)
	conn.Close()
}

func RetrieveKeys(prefix string) string {
	return "{}"
}
