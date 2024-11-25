FROM golang:1.23.3
WORKDIR /app
COPY . .
RUN go build ./cmd/swa
CMD ["./swa"]
