FROM golang:alpine

RUN apk update && \
    apk add --no-cache \
        make

COPY . /app
WORKDIR /app

ENV HOME='/root'
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN go mod vendor && \
    go mod tidy

RUN make build

CMD ["make", "run"]

