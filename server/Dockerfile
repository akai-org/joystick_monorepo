FROM golang:1.17

WORKDIR /app
COPY . .
RUN go mod download && go build -o $GOPATH/bin/app main.go

CMD ["app"]