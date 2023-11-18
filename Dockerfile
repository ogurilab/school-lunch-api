FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app

COPY . .
RUN go build -o main app/cmd/main.go


FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY run.sh .
COPY wait-for.sh .
RUN chmod +x run.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/run.sh" ]