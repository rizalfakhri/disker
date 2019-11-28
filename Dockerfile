FROM golang:1.11

LABEL maintainer="Rizal Fakhri <me@rizalfakhri.id>"

WORKDIR /go/src/github.com/rizalfakhri/disker

COPY . .

COPY wait-for-it.sh /wait-for-it.sh

RUN chmod +x /wait-for-it.sh

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["disker", "-channel=rabbitmq", "-data=disk"]
