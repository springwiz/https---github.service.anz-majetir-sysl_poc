ARG REGISTRY

FROM ${REGISTRY}${REGISTRY:+/}golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o addressservice ./cmd/addressservice

FROM ${REGISTRY}${REGISTRY:+/}alpine
WORKDIR /app
COPY --from=builder /app/addressservice /bin/
CMD addressservice
