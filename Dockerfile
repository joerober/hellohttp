FROM golang AS builder
WORKDIR /opt
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o hellohttp .

FROM golang
WORKDIR /opt
COPY --from=0 /opt/hellohttp .
CMD ["/opt/hellohttp"]