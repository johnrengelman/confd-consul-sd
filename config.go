// Copyright (c) 2014 John Engelman. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package main

import (
	"flag"
)

var (
	config       Config
	debug        bool
	quiet        bool
	verbose      bool
	transport    string
	port         int
	socket       string
	printVersion bool
)

// A Config structure is used to configure confd-consul-sd.
type Config struct {
	Debug     bool
	Quiet     bool
	Verbose   bool
	Transport string
	Port      int
	Socket    string
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

// processFlags iterates through each flag set on the command line and
// overrides corresponding configuration settings.
func processFlags() {
	flag.Visit(setConfigFromFlag)
}

func setConfigFromFlag(f *flag.Flag) {
	switch f.Name {
	case "debug":
		config.Debug = debug
	case "quiet":
		config.Quiet = quiet
	case "verbose":
		config.Verbose = verbose
	case "transport":
		config.Transport = transport
	case "port":
		config.Port = port
	case "socket":
		config.Socket = socket
	}
}
