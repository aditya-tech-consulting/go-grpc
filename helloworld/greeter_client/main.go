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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName       = "world"
	defaultSubject    = "Subject"
	defultId          = "Id"
	postProfessor     = "postProfessor"
	postCourse        = "postCourse"
	postStudent       = "postStudent"
	postStudentCourse = "postStudentCourse"
	getStudentCourse  = "getStudentCourse"
	defualtMethodName = "No Method"
)

var (
	addr          = flag.String("addr", "localhost:50051", "the address to connect to")
	name          = flag.String("name", defaultName, "Name to greet")
	subject       = flag.String("subject", defaultSubject, "Subject to Teach")
	id            = flag.String("id", defultId, "id to create")
	courseName    = flag.String("courseName", "Physics", "A Course Name")
	studentName   = flag.String("studentName", "None", "A Student's Name")
	professorName = flag.String("professorName", "Dr. Watson", "The name of a professor")
	methodName    = flag.String("methodName", defualtMethodName, "Method Name to Call")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Input is : %s", *methodName)
	log.Printf("Name: %s", *name)
	if *methodName == postProfessor {
		stream, err := c.CreateProfessor(ctx, &pb.ProfessorRequest{Name: *name, Subject: *subject, Id: *id})
		if err != nil {
			log.Fatalf("could not create professor: %v", err)
		}
		log.Printf("Professor Name: %s, Subject Name: %s", stream.GetName(), stream.GetSubject())

	} else if *methodName == postCourse {
		streams, err := c.CreateCourse(ctx, &pb.CourseRequest{Name: *name, Id: *id})
		if err != nil {
			log.Fatalf("could not create course: %v", err)
		}
		log.Printf("Course Name: %s", streams.GetName())
	} else if *methodName == postStudent {
		streams, err := c.CreateStudent(ctx, &pb.StudentRequest{Name: *studentName, Id: *id})
		if err != nil {
			log.Fatalf("could not create student: %v", err)
		}
		log.Printf("Student Name: %s", streams.GetName())
	} else if *methodName == getStudentCourse {
		streams, err := c.GetStudentCourse(ctx, &pb.StudentCourseSearchRequest{StudentName: *studentName, CourseName: *courseName})
		if err != nil {
			log.Fatalf("could not retrive student details: %v", err)
		}
		log.Printf("Course Name: %s, Student Name:%s, Professor Name: %s", streams.GetCourseName(), streams.GetCourseName(), streams.GetProfessorName())
	} else if *methodName == getStudentCourse {
		streams, err := c.GetStudentCourse(ctx, &pb.StudentCourseSearchRequest{StudentName: *studentName, CourseName: *courseName})
		if err != nil {
			log.Fatalf("could not retrive student details: %v", err)
		}
		log.Printf("Course Name: %s, Student Name:%s, Professor Name: %s", streams.GetCourseName(), streams.GetCourseName(), streams.GetProfessorName())
	} else if *methodName == postStudentCourse {
		streams, err := c.CreateStudentCourse(ctx, &pb.StudentCourseRequest{StudentName: *studentName, CourseName: *courseName, ProfessorName: *professorName})
		if err != nil {
			log.Fatalf("could not create student and course details: %v", err)
		}
		log.Printf("Course Name: %s, Student Name:%s, Professor Name: %s", streams.GetCourseName(), streams.GetStudentName(), streams.GetProfessorName())
	} else if *methodName == getStudentCourse {
		streams, err := c.GetStudentCourse(ctx, &pb.StudentCourseSearchRequest{StudentName: *studentName, CourseName: *courseName})
		if err != nil {
			log.Fatalf("could not retrive student details: %v", err)
		}
		log.Printf("Course Name: %s, Student Name:%s, Professor Name: %s", streams.GetCourseName(), streams.GetCourseName(), streams.GetProfessorName())
	} else {
		log.Print("Hello \nYour Input did not match \nNo Worries\n Try again ")
	}

	/*r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())*/

}
