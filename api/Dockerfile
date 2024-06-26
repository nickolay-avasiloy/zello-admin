FROM golang:1.22.4

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv

# Go app
RUN go build -o /go/bin/api

EXPOSE 8080

CMD ["/go/bin/api"]
