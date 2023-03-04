FROM golang:alpine as builder

WORKDIR /work

ADD . /work

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM amd64/alpine:latest

WORKDIR /work

COPY --from=builder /work/main /work/main
ADD ./contract/IAllERC20.abi /work/contract/IAllERC20.abi

CMD ["./main"]
