package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"session-3-todos/common/config"
	"session-3-todos/common/models"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var localStorage *models.TodoList

func init() {
	localStorage = new(models.TodoList)
	localStorage.List = make([]*models.Todo, 0)
}

type TodosServer struct {
	models.UnimplementedTodosServer
}

func (TodosServer) Add(ctx context.Context, param *models.Todo) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Add Activity - ", param.String())
	return new(empty.Empty), nil
}

func (TodosServer) List(ctx context.Context, void *empty.Empty) (*models.TodoList, error) {
	return localStorage, nil
}

func (TodosServer) Delete(ctx context.Context, deleteByID *models.TodoId) (*empty.Empty, error) {
	for index, data := range localStorage.List {
		if data.Id == deleteByID.Id {
			localStorage.List = append(localStorage.List[:index], localStorage.List[index+1:]...)
		}
	}
	return new(empty.Empty), nil
}

func (TodosServer) Update(ctx context.Context, reqUpdate *models.TodoUpdate) (*empty.Empty, error) {
	for index, data := range localStorage.List {
		if data.Id == reqUpdate.Id {
			localStorage.List[index].Activity = reqUpdate.Activity
		}
	}
	return new(empty.Empty), nil
}

func main() {
	srv := grpc.NewServer()
	todoSrv := TodosServer{}
	models.RegisterTodosServer(srv, &todoSrv)

	log.Println("Starting User Server at ", config.SERVICE_TODO_PORT)

	l, err := net.Listen("tcp", config.SERVICE_TODO_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_TODO_PORT, err)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_TODO_PORT, "gRPC server endpoint")
		_ = models.RegisterTodosHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting User Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()
	log.Fatalln(srv.Serve(l))
}
