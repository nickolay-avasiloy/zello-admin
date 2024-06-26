FROM golang:1.22.4

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

