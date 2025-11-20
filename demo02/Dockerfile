FROM golang:1.25.4-alpine3.22 AS step1
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o demo main.go


FROM alpine:3.22
WORKDIR /app
COPY --from=step1 /app/demo .

EXPOSE 1323
CMD ["./demo"]