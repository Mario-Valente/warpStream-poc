FROM golang:1.24.2-alpine3.20 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build  -o /go/bin/app

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

COPY --from=builder /go/bin/app /app

CMD ["/app"]
