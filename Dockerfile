From golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV GO111MODULE=on

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o SignUpPage

CMD ["./SignUpPage"]