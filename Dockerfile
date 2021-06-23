FROM golang:1.16 as builder

WORKDIR /workdir

COPY go.mod go.mod
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o hello main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workdir/hello .
USER 65532:65532

ENTRYPOINT ["/hello"]

