FROM golang:alpine as builder
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN export CGO_ENABLED=0
RUN go build -o server cmd/main.go

FROM alpine
WORKDIR /src
COPY --from=builder /src/server /src/server 
COPY --from=builder /src/template /src/template 
CMD ["./server"]