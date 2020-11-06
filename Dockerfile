FROM golang:1.14

WORKDIR /app

# had issues installing with GOPATH and it can't be set with go.mod
ENV GOPATH ''

COPY . .

RUN go install
