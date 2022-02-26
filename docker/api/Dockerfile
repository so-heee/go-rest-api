FROM golang:1.16-alpine

WORKDIR /go/api
COPY ./api .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
