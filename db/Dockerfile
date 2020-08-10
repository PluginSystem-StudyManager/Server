FROM yottadb/yottadb-base

RUN apt update
RUN apt install -y git 

WORKDIR /db
ENV GOBIN /db/bin
ENV GOPATH /db
ENV PATH=$PATH:$GOPATH/bin

COPY . .

RUN apt install -y curl
RUN wget -q https://golang.org/dl/go1.14.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.7.linux-amd64.tar.gz; 

ENV PATH=/usr/local/go/bin:$PATH
ENV PATH=/db:$PATH

RUN go get lang.yottadb.com/go/yottadb

RUN go get -v
RUN go install -v

EXPOSE 8090

RUN chmod +x start

CMD ["start"]

