FROM golang:1.24.0

EXPOSE 8010

WORKDIR /app

# Protocol Buffers (must be compiled)
COPY proto/gen /proto/gen
# Character service
COPY services/dnd/ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /dnd-server ./exec/grpc 

CMD /dnd-server -character_service_addr $CHARACTERS_SERVICE_ADDR