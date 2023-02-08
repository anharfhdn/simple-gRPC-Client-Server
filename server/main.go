package main

import (
	pb "anharfhdn/learn/simple-grpc/student"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	fmt.Println("incoming request...")

	for _, value := range d.students {
		if value.Email == student.Email {
			return value, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/students.json")
	if err != nil {
		log.Fatalln("Error in read file", err.Error())
	}
	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("Error in unmarshal data json", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("Error in listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error when serve grpc", err.Error())

	}

}
