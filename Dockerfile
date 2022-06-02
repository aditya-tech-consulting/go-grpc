FROM golang:1.16-alpine
#FROM busybox

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
#COPY go.sum .
COPY  ./helloworld ./
#COPY go.mod /
#COPY go.sum /

RUN go get google.golang.org/grpc

RUN apk update
RUN apk add git
#RUN go clean -modcache
#RUN git clone https://github.com/aditya-tech-consulting/go-grpc
#RUN cd go-grpc
#RUN go install -tags "v1.0.1" github.com/aditya-tech-consulting/go-grpc/helloworld/helloworld@master
RUN go get github.com/aditya-tech-consulting/go-grpc/helloworld/helloworld
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
EXPOSE 8080

# Run
ENTRYPOINT [ "go"]
CMD [ "run","./greeter_server/main.go"]
