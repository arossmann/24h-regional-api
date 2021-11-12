FROM golang:1.17 AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM scratch

COPY --from=builder /main ./
ENTRYPOINT ["./main"]
EXPOSE 8080