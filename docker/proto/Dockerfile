FROM golang:1.22.6

COPY proto/ ./proto

RUN touch proto/test.txt

# Install protoc
RUN apt update && \
    apt upgrade -y && \
    apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN mkdir -p proto/gen

# characters
RUN protoc \
    --proto_path=proto \
    --go_out=./proto/gen \
    --go_opt=paths=source_relative \
    --go_opt=Mcharacters/service.proto=proto/characters \
    --go-grpc_out=./proto/gen \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=Mcharacters/service.proto=proto/characters \
    characters/service.proto

# Turn it into a Go Module
RUN cd proto/gen && \
    go mod init github.com/AranGarcia/proto/gen
