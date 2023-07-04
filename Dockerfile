FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o sellerApp

EXPOSE 8000
CMD ["./sellerApp"]