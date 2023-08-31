FROM golang:latest

WORKDIR /opt/app
COPY ./app/server /opt/app

RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 80

CMD ["go", "run", "main.go"]