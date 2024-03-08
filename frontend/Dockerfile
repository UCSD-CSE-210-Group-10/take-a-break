# Base image with Ubuntu
FROM ubuntu:latest

# Install package managers
RUN apt-get update && apt-get install -y git curl build-essential ca-certificates software-properties-common

# Install Node Version Manager (NVM)
RUN curl -sL https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
RUN . ~/.bashrc && nvm install 20.11

# Set working directory (project root)
WORKDIR /app

COPY package.json .
RUN . ~/.bashrc && npm install
COPY . ./

EXPOSE 3000:3000
EXPOSE 8080:8080

CMD . ~/.bashrc && npm start