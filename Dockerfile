FROM golang:stretch as builder
WORKDIR /go
COPY . .
RUN CGO_ENABLED=0 go install hello

FROM scratch 
COPY --from=builder /go/bin/hello /usr/local/hello
EXPOSE 8000
ENTRYPOINT ["/usr/local/hello"]