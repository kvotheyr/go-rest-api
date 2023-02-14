FROM golang:1-buster
ENV CGO_ENABLED 0
ENV PACKAGE_DIR '.'

RUN go install github.com/go-delve/delve/cmd/dlv@latest && go install github.com/cosmtrek/air@latest

ENV GOFLAGS '-mod=vendor'

WORKDIR /app
COPY ./start.sh /start.sh
COPY ./.air.toml /.air.toml
EXPOSE 40000

ENTRYPOINT ["sh", "/start.sh"]