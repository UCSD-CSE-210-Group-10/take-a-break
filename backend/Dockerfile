FROM golang:1.22.0-alpine

WORKDIR /app

COPY . .

# Install Go dependencies
RUN go mod download

EXPOSE 8080

COPY config.sh .

# PostgreSQL Configuration
ENV DB_HOST=172.18.0.2
ENV DB_PORT=5432
ENV DB_NAME=take-a-break
ENV DB_USER=postgres
ENV DB_PASSWORD=123456

# OAuth Configuration
ENV GOOGLE_CLIENT_ID=256148397214-ivsp4a1ro4posdhv4iiot28mb3gb5s7n.apps.googleusercontent.com
ENV GOOGLE_CLIENT_SECRET=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
ENV AUTHURL=https://accounts.google.com/o/oauth2/v2/auth
ENV TOKENURL=https://oauth2.googleapis.com/token
ENV POSTURL=https://jsonplaceholder.typicode.com/posts
ENV REDIRECT_URL=http://35.192.5.163.nip.io:3000/
ENV CLIENT_URL=http://35.192.5.163.nip.io:3000
ENV TOKEN_SECRET=123456

CMD go run main.go 
