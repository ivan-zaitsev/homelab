FROM golang:latest

WORKDIR /app
COPY . /app

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT [ "owntracks-nextcloud-proxy" ]

EXPOSE 8000
