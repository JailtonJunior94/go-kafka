FROM golang:1.15.7-alpine

WORKDIR /go/src/

COPY . .

RUN ls
RUN go clean --modcache
RUN GOOS=linux go build cmd/api/main.go

EXPOSE 3000
ENTRYPOINT ["./main"]