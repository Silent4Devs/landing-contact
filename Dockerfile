FROM golang:1-alpine

# Install MySQL client
RUN apk add --no-cache mysql-client

WORKDIR /app
# COPY go.mod .
# COPY go.sum .
COPY . .

RUN go mod tidy && go build .

ENTRYPOINT ["./fiber-boilerplate"]