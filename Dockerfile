FROM alpine:latest

WORKDIR /work

ADD ./bin/coin-manage /work/main

CMD ["./main"]

