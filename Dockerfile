FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN ls
RUN go mod tidy
RUN go test ./test -v


# # Stage 2: Create the final image
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /app

# RUN go build -o client main/main.go

# # Copy the Go binary from the build stage
# COPY --from=build /app/client .

# # Start the client
# CMD ["./client"]
