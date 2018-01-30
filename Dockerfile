FROM golang

WORKDIR /go/src/github.com/WeltN24/gossad2
COPY . .

RUN go get -d -v ./...
RUN go build -v -o gossad2
EXPOSE 5000

CMD ["./gossad2"]