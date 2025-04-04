FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go build -o whatwhen .

RUN chmod +x whatwhen

ENTRYPOINT ["./whatwhen"]
