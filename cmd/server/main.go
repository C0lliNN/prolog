package main

import (
	"flag"
	"fmt"
	"github.com/c0llinn/prolog/internal/log"
	"github.com/c0llinn/prolog/internal/server"
	defaultlog "log"
	"net"
)

var (
	port          = flag.Int("port", 8080, "Specifies the port the server will run")
	dir           = flag.String("dir", "", "Directory path the log files will be stored")
	maxStoreBytes = flag.Uint64("max-store-bytes", 1024, "Max Store Bytes for each Log Segment")
	maxIndexBytes = flag.Uint64("max-index-bytes", 1024, "Max Index Bytes for each Log Segment")
	initialOffset = flag.Uint64("initial-offset", 0, "Initial Offset")
)

func main() {
	flag.Parse()

	config := createLogConfig()

	l, err := log.NewLog(*dir, config)
	if err != nil {
		defaultlog.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		defaultlog.Fatal(err)
	}

	srv, err := server.NewGRPCServer(&server.Config{CommitLog: l})
	if err != nil {
		defaultlog.Fatal(err)
	}

	defaultlog.Print("gRPC server starting at port ", *port)
	defaultlog.Fatal(srv.Serve(lis))
}

func createLogConfig() log.Config {
	return log.Config{
		Segment: struct {
			MaxStoreBytes uint64
			MaxIndexBytes uint64
			InitialOffset uint64
		}{MaxStoreBytes: *maxStoreBytes, MaxIndexBytes: *maxIndexBytes, InitialOffset: *initialOffset},
	}
}
