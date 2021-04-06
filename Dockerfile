FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app/cmd/test
RUN go build -o main .
CMD ["/app/cmd/test/main"]