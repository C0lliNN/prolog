package main

import (
	"context"
	"flag"
	"fmt"
	api "github.com/c0llinn/prolog/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	uri    = flag.String("uri", "localhost:8080", "The address the client will connect to")
	client api.LogClient
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client = api.NewLogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	executeAction(ctx, os.Args)
}

func executeAction(ctx context.Context, args []string) {
	if len(args) < 3 {
		log.Fatal("You need to provide a command [append, read] and and argument ['value', offset]")
	}

	switch args[1] {
	case "append":
		appendRecord(ctx, []byte(os.Args[2]))
	case "read":
		offset, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("The offset must be an integer")
		}

		readRecord(ctx, uint64(offset))
	default:
		log.Fatal("The provided command is not implemented.")
	}
}

func appendRecord(ctx context.Context, value []byte) {
	res, err := client.Produce(ctx, &api.ProduceRequest{Record: &api.Record{Value: value}})
	if err != nil {
		log.Fatalf("could not insert a log record: %v", err)
	}

	fmt.Println("Record was inserted successfully with offset: ", res.Offset)
}

func readRecord(ctx context.Context, offset uint64) {
	res, err := client.Consume(ctx, &api.ConsumeRequest{Offset: offset})
	if err != nil {
		log.Fatalf("could not consume record: %v", err)
	}

	fmt.Println(string(res.Record.Value))
}