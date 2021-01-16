FROM golang:1.15.6-alpine

COPY . /app

WORKDIR /app
RUN go build

CMD ./hospital
