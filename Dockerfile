# Use the official Golang image as the base image for Windows
# You can specify the version you need. For example, "golang:1.16"
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main

# Expose the port that your Go API will listen on
EXPOSE 8084

# Define the command to run your Go API when the container starts
CMD ["./main"]

# Commands create
#  docker buildx build --platform linux/arm/v84 -t articlesfeedapi . 
#  docker buildx build -t articlesfeedapi . 
# Commands install
#  docker run -p 8084:8084 -e CONNECTIONSTRING='xxxxxx' articlesfeedapi
