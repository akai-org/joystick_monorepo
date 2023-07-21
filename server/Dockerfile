FROM golang:1.18

WORKDIR /app
COPY . .
RUN go mod download && go build -o $GOPATH/bin/app

CMD ["app"]