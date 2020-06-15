FROM golang AS build
WORKDIR /workspace
COPY go.mod .
RUN go mod download -json
COPY hello-server.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -i -installsuffix cgo -o hello-server

FROM scratch
COPY --from=build /workspace/hello-server /
USER 1001
ENTRYPOINT ["/hello-server"]
