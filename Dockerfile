FROM golang:1.20-alpine

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@v1.29.0
WORKDIR /app 

CMD ["air", "-c", ".air.toml"]
