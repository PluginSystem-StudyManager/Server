FROM golang:1.14

WORKDIR /server
ENV GOPATH /server
ENV PATH=$PATH:$GOPATH/bin

COPY . .

WORKDIR /server/src
RUN go get -d ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["server"]
