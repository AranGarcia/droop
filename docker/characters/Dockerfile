FROM golang:1.24.0

EXPOSE 8090

WORKDIR /app

# Protocol Buffers (must be compiled)
COPY proto/gen /proto/gen
# Shared libraries
COPY shared/ /shared
# Character service
COPY services/characters/ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /characters-server ./exec/grpc 

CMD /characters-server -mongo_host $MONGO_HOST