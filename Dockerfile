FROM golang:alpine as builder

ENV GO111MODULE=off

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go get github.com/go-chi/chi

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

COPY --from=builder /app/bin/main .

CMD ["./main"]