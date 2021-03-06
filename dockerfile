FROM golang:1.15

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "main.go"]