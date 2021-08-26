FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .

EXPOSE 8080
CMD ["/todo-list"]