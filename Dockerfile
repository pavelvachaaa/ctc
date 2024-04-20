FROM golang:latest

LABEL maintainer="Pavel Vacha <pavel.vacha@tul.cz>"

WORKDIR /apk

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build app/main.go -o main .

EXPOSE 8080

CMD ["./main"]