FROM golang:1.17 as builder

WORKDIR /workspace
COPY go.* .
RUN go mod download
COPY *.go .
COPY templates templates
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /http-debug .


# Use distroless as minimal base image to package the project
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /http-debug .
ENTRYPOINT ["/http-debug"]

