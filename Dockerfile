FROM golang:1.16-alpine
ADD . /go/src/github.com/toDo-list
WORKDIR /go/src/github.com/toDo-list
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /todo-list

EXPOSE 8080
CMD ["/todo-list"]