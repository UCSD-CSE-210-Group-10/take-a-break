FROM golang:1.22.0-alpine

WORKDIR /app

COPY . .

# Install Go dependencies
RUN go mod download

EXPOSE 8080

COPY config.sh .

RUN chmod +x config.sh && ./config.sh

RUN echo $DB_PORT

CMD go run main.go 
