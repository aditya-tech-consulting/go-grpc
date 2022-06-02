/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "github.com/aditya-tech-consulting/go-grpc/helloworld/helloworld"
)

const (
	ERR_LOG  = "error.log"
	INFO_LOG = "info.log"
)

var (
	port           = flag.Int("port", 8080, "The server port")
	students       = make(map[string]Student)
	professors     = make(map[string]Professor)
	courses        = make(map[string]Course)
	studentCourses = make(map[string][]StudentCourse)
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

type Professor struct {
	Name    string
	Subject string
	Id      string
}

type Course struct {
	Name string
	Id   string
}

type Student struct {
	Id   string
	Name string
}

type StudentCourse struct {
	studentName   string
	courseName    string
	professorName string
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fileInfo, err := openLogFile(INFO_LOG)
	log.Print("Error Generated:", err)
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
func (s *server) CreateProfessor(ctx context.Context, in *pb.ProfessorRequest) (*pb.ProfessorReply, error) {
	p1 := Professor{in.GetName(), in.GetSubject(), in.GetId()}
	professors[in.GetId()] = p1
	fileInfo, err := openLogFile(INFO_LOG)
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	fmt.Println("Professors:", professors)
	infoLog.Printf("Professors:", professors)
	log.Print("Error Generated:", err)
	return &pb.ProfessorReply{Name: in.GetName(), Subject: in.GetSubject(), Id: in.GetId()}, nil

}

func (s *server) CreateCourse(ctx context.Context, in *pb.CourseRequest) (*pb.CourseReply, error) {
	c1 := Course{in.GetName(), in.GetId()}

	courses[in.GetId()] = c1
	fileInfo, err := openLogFile(INFO_LOG)
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	log.Print("Error Generated:", err)
	fmt.Println("Courses:", courses)
	infoLog.Printf("Courses:", courses)
	return &pb.CourseReply{Name: in.GetName(), Id: in.GetId()}, nil

}

func (s *server) CreateStudent(ctx context.Context, in *pb.StudentRequest) (*pb.StudentReply, error) {

	student := Student{in.GetName(), in.GetId()}
	students[in.GetId()] = student
	fileInfo, err := openLogFile(INFO_LOG)
	log.Print("Error Generated:", err)
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	fmt.Println("Students:", students)
	infoLog.Printf("Students:", students)
	return &pb.StudentReply{Name: in.GetName(), Id: in.GetId()}, nil

}
func (s *server) GetStudentCourse(ctx context.Context, in *pb.StudentCourseSearchRequest) (*pb.StudentCourseSearchReply, error) {
	//studentCourse := StudentCourse{in.GetStudentName(), in.GetCourseName(), in.GetProfessorName()}
	studentCourseList := studentCourses[in.GetCourseName()]
	var studentCourseTemp StudentCourse
	for _, studentCourseData := range studentCourseList {
		if studentCourseData.courseName == in.GetCourseName() {
			studentCourseTemp := studentCourseData
			fmt.Print(studentCourseTemp)
			break
		}
	}
	return &pb.StudentCourseSearchReply{StudentName: studentCourseTemp.studentName, CourseName: studentCourseTemp.courseName, ProfessorName: studentCourseTemp.professorName}, nil
}
func (s *server) CreateStudentCourse(ctx context.Context, in *pb.StudentCourseRequest) (*pb.StudentCourseReply, error) {
	studentCourse := StudentCourse{in.GetStudentName(), in.GetCourseName(), in.GetProfessorName()}
	studentCourseList := studentCourses[in.GetCourseName()]
	studentCourseList = append(studentCourseList, studentCourse)
	studentCourses[in.GetCourseName()] = studentCourseList
	fileInfo, err := openLogFile(INFO_LOG)
	log.Print("Error Generated:", err)
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	fmt.Println("Student Course List:", studentCourses)
	infoLog.Printf("Student Course List:", studentCourses)
	return &pb.StudentCourseReply{StudentName: in.GetStudentName(), CourseName: in.GetCourseName(), ProfessorName: in.GetProfessorName()}, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
