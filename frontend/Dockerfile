# Base image with Ubuntu
FROM ubuntu:latest

# Install package managers
RUN apt-get update && apt-get install -y git curl build-essential ca-certificates software-properties-common

ENV GOOGLE_CLIENT_ID=256148397214-ivsp4a1ro4posdhv4iiot28mb3gb5s7n.apps.googleusercontent.com
ENV GOOGLE_CLIENT_SECRET=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
ENV AUTHURL=https://accounts.google.com/o/oauth2/v2/auth
ENV TOKENURL=https://oauth2.googleapis.com/token
ENV POSTURL=https://jsonplaceholder.typicode.com/posts
ENV REDIRECT_URL=http://34.172.41.112.nip.io:3000/
ENV CLIENT_URL=http://34.172.41.112.nip.io:3000
ENV TOKEN_SECRET=123456

# OAuth Configuration - Frontend (React App)
ENV REACT_APP_SERVER_URL=http://34.172.41.112:8080
ENV REACT_APP_GOOGLE_CLIENT_ID=256148397214-ivsp4a1ro4posdhv4iiot28mb3gb5s7n.apps.googleusercontent.com
ENV REACT_APP_REDIRECT_URL=http://34.172.41.112.nip.io:3000/
ENV REACT_APP_CLIENT_URL=http://34.172.41.112.nip.io:3000
ENV REACT_APP_AUTHURL=https://accounts.google.com/o/oauth2/v2/auth
ENV REACT_APP_TOKENURL=https://oauth2.googleapis.com/token
ENV REACT_APP_POSTURL=https://jsonplaceholder.typicode.com/posts

# Install Node Version Manager (NVM)
RUN curl -sL https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
RUN . ~/.bashrc && nvm install 20.11

# Set working directory (project root)
WORKDIR /app

COPY package.json .
RUN . ~/.bashrc && npm install
COPY . .

EXPOSE 3000

CMD . ~/.bashrc && npm start