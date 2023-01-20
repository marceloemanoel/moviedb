FROM golang:1.19 as dev

RUN apt update && apt install -y netcat
RUN go install github.com/githubnemo/CompileDaemon@latest
ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

FROM golang:1.19-alpine as build

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o movieDB main.go


FROM alpine as runtime

COPY --from=build /app/wait-for-it.sh /app/
COPY --from=build /app/movieDB /app/movieDB

CMD ["/app/movieDB"]