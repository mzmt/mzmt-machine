FROM golang:latest

WORKDIR /go/src
ADD . /go/src

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3
RUN go get strconv

CMD ["go", "run", "main.go"]
