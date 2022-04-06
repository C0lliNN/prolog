package main

import (
	"flag"
	"fmt"
	"github.com/c0llinn/prolog/internal/auth"
	"github.com/c0llinn/prolog/internal/config"
	"github.com/c0llinn/prolog/internal/log"
	"github.com/c0llinn/prolog/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	logger, _ := zap.NewDevelopment()
	logger = logger.Named("server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal(err.Error())
	}

	logConfig := createLogConfig()

	l, err := log.NewLog(*dir, logConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	serverConfig := &server.Config{
		CommitLog:  l,
		Authorizer: auth.New(config.ACLModelFile, config.ACLPolicyFile),
	}
	serverTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile: config.ServerCertFile,
		KeyFile:  config.ServerKeyFile,
		CAFile:   config.CAFile,
		Server:   true,
	})

	srv, err := server.NewGRPCServer(serverConfig, grpc.Creds(credentials.NewTLS(serverTLSConfig)))
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info(fmt.Sprintf("gRPC server starting at port %d", *port))
	logger.Fatal(srv.Serve(lis).Error())
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
