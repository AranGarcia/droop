FROM golang:1.23.2

EXPOSE 8090

WORKDIR /app

# Protocol Buffers (must be compiled)
COPY proto/gen /proto/gen
# Character service
COPY characters/ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /characters-server ./exec/grpc 

CMD /characters-server -mongo_host $MONGO_HOST