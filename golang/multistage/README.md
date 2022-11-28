# GoLang Multi-stage Dockerfile Example

Build:
docker build -t multistage .

Run:
docker run --rm -it  -p 8080:8080/tcp multistage:latest
