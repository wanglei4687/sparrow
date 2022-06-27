FROM golang:1.18 as builder

COPY go.mod go.mod
COPY go.sum go.sum

RUN go download

RUN CGO_ENABLED=0 go build -a -o bin/sparrow cmd/sparrow/main.go

FROM alpine:3.16

COPY --from=builder /bin/sparrow .

ENTRYPOINT ["/sparrow"]
