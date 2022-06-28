FROM golang:1.16-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .


FROM alpine:latest AS production
COPY  --from=builder /app .
EXPOSE 8080
CMD [ "./main" ]