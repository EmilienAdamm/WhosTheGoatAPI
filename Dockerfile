LABEL authors="Emilien"

#---------BUILDING---------#
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#---------RUNNING---------#
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]