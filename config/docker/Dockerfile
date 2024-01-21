FROM golang:1.21

RUN go get -u github.com/lib/pq

WORKDIR /TestTaskGo

COPY . .

RUN go build -o main .

ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=postgres
ENV DB_SSL_MODE=disable

CMD ["./main"]
