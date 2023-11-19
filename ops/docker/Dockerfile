FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app

COPY . .
RUN go build -o main app/cmd/main.go

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/main .
COPY app/.env .
COPY ops/scripts/*.sh .
RUN chmod +x *.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/run.sh" ]
