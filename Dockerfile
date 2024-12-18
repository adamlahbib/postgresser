FROM golang:alpine as builder
RUN apk --no-cache --update add ca-certificates make git 
WORKDIR /app
COPY . .
RUN make build

FROM scratch 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/main /app 
CMD ["/app"]