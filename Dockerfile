FROM golang:1.21

RUN mkdir -p /app

ADD /app/. /app/

WORKDIR /app

RUN go build .

CMD ["./go-batch"] 
