FROM yottadb/yottadb AS yottadb

FROM golang

RUN apt update ; apt install libtinfo5
COPY --from=yottadb /opt/yottadb/current /opt/yottadb/current
ENV PKG_CONFIG_PATH=/opt/yottadb/current

WORKDIR /db
ENV GOBIN /db/bin
ENV GOPATH /db/go
ENV PATH=$PATH:/db/bin

COPY . .

RUN . /opt/yottadb/current/ydb_env_set ; go get lang.yottadb.com/go/yottadb


RUN go get -v
RUN go install -v

EXPOSE 8090

RUN chmod +x start

CMD ["start"]

