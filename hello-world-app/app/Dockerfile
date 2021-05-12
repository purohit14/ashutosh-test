FROM golang:1.14

WORKDIR /go/src/app
COPY . .
COPY ../data .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]