FROM golang:alpine as builder

WORKDIR /go/src/
COPY . . 

ENV GO111MODULE=on GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux

RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/emby_exporter

# Binary container
FROM scratch
COPY --from=builder /go/bin/emby_exporter /bin/emby_exporter

ENTRYPOINT ["/bin/emby_exporter"]