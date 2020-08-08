FROM alpine:latest

RUN apk add --no-cache git go npm

WORKDIR /server
ENV GOPATH /server
ENV PATH=$PATH:$GOPATH/bin

COPY . .

RUN npm install
RUN npm run grunt

WORKDIR /server/src
RUN go get -d ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["server"]
