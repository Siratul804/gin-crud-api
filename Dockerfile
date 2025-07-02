# Base image: Ubuntu with Go & Git
FROM ubuntu:22.04

# Install Git and Go
RUN apt-get update && \
    apt-get install -y --no-install-recommends git golang-go && \
    rm -rf /var/lib/apt/lists/*

# Set working dir
WORKDIR /app

# Clone this repo (replace URL when pushing to your GitHub)
ARG REPO_URL=https://github.com/Siratul804/gin-crud-api
RUN git clone "$REPO_URL" .

# Build the binary
RUN go build -o server main.go

# Expose port
EXPOSE 8080
ENV PORT=8080

# Run
CMD ["/app/server"]