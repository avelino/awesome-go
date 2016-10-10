FROM golang:alpine

RUN apk add --update -t build-deps curl go git libc-dev gcc libgcc
RUN go get github.com/russross/blackfriday github.com/gorilla/mux

WORKDIR /srv

CMD ["go", "run", "repo.go"]
