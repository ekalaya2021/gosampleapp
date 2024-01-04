
FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8081
RUN go build -o main .
CMD ["/app/main"]
