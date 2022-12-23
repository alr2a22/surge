ARG REGISTRY

# build stage
FROM ${REGISTRY}/library/golang:1.19.4 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /main ./main.go

# final stage
FROM ${REGISTRY}/library/alpine:3.17

RUN mkdir /app

COPY --from=builder /main /app/main
EXPOSE 3000
