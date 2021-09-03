FROM golang:1.16

RUN mkdir /ops
WORKDIR /ops

COPY ./ /ops

RUN go build

EXPOSE 8081

CMD ["./ops"]