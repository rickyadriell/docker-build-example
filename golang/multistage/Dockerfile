# Build stage
FROM golang:1.19.3-alpine3.16 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# # # Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=build /app .
EXPOSE 8080
CMD ["./main"]
