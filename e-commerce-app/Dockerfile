FROM golang:1.22.2-alpine

WORKDIR /app

COPY . .

RUN go build -o e-commerce-app .

CMD ["./e-commerce-app"]
