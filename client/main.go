package main

import (
	"context"
	"fmt"
	"log"
	"session-3-todos/common/config"
	"session-3-todos/common/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serviceTodo() models.TodosClient {
	port := config.SERVICE_TODO_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return models.NewTodosClient(conn)
}

func main() {
	todoSvc := serviceTodo()
	fmt.Println("todo test")
	var todo1 = &models.Todo{
		Id:       "u001",
		Activity: "Nabung di bank",
	}

	todoSvc.Add(context.Background(), todo1)

	// todos, _ := todoSvc.List(context.Background(), new(empty.Empty))
	// log.Printf("List Users %+v\n", todos.GetList())
}
