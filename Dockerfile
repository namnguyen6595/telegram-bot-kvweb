FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server .
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]