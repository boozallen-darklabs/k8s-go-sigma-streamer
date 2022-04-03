# Warning: Invoke build from docker-compose.yaml
FROM golang:1.16.0 AS build
RUN useradd -u 10001 benthos
WORKDIR /build
ENV GO111MODULE=on
COPY engine/go.mod engine/go.sum ./
COPY engine/processor/processor.go ./processor/
RUN go mod download

COPY engine/main.go . 
RUN CGO_ENABLED=0 GOOS=linux go build -o singe main.go

FROM busybox AS package
WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /build/singe .
COPY engine/rules rules
COPY engine/benthos.yaml .
USER benthos
EXPOSE 4195
ENTRYPOINT ["./singe"]
CMD ["-c", "benthos.yaml"]
