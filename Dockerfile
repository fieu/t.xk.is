FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /main

FROM scratch

WORKDIR /

COPY --from=builder /main /main

ENTRYPOINT ["/main"]
