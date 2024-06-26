FROM golang:1.22.4

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u github.com/golang-migrate/migrate/v4


EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

