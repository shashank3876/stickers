FROM golang:latest

WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Set environment variables

COPY config.json /app/config.json





EXPOSE 9090

CMD ["go", "run", "app/main.go"]

