FROM golang:1.21 AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM alpine:3.19.0
CMD apk update && apk add bind-tools
COPY --from=builder /main ./
ENV MONGODB_USERNAME=MONGODB_USERNAME MONGODB_PASSWORD=MONGODB_PASSWORD MONGODB_ENDPOINT=MONGODB_ENDPOINT MONGODB_DATABASE=MONGODB_DATABASE MONGODB_COLLECTION=MONGODB_COLLECTION PORT=PORT
ENTRYPOINT ["./main"]
EXPOSE 8888