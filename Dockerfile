FROM golang:1.16

COPY ./binary /go/src 
EXPOSE 8081

CMD ["./src/binary"]