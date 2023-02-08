package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "anharfhdn/learn/simple-grpc/student"

	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}

	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("Error when get student by email", err.Error())
	}
	fmt.Println(student)

}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	connection, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("Error in Dial", err.Error())
	}

	defer connection.Close()

	client := pb.NewDataStudentClient(connection)
	getDataStudentByEmail(client, "youremail@gmail.com")
	getDataStudentByEmail(client, "hername@gmail.com")
}
