FROM golang:latest

WORKDIR /opt/app
COPY ./app/server /opt/app

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
ENV PATH $PATH:/go/bin/linux_amd64

EXPOSE 80

CMD ["air", "-c", ".air.toml"]