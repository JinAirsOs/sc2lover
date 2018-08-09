FROM golang:alpine as builder
RUN apk add --no-cache ca-certificates
WORKDIR /go/src/github.com/JinAirsOs/sc2lover
COPY . .
RUN go generate ./static
RUN CGO_ENABLED=0 go install ./cmd/sc2lover

FROM alpine
WORKDIR /app
COPY --from=builder /go/bin/sc2lover /app/sc2lover
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["/app/sc2lover"]
