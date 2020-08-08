FROM alpine:latest

RUN apk add --no-cache git go npm
RUN go get github.com/Joker/jade/cmd/jade

WORKDIR /server
ENV GOPATH /server
ENV GOROOT /bin
ENV PATH=$PATH:$GOPATH/bin

COPY . .

#RUN npm install
#RUN npm run grunt

RUN printenv

WORKDIR /server/src
RUN go get -d ./...
RUN go generate
RUN go install -v ./...


EXPOSE 8080

CMD ["server"]
