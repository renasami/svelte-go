FROM golang:latest
WORKDIR /go/src
COPY ./src .
# 以下追記
# 以下編集
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
