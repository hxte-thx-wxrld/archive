FROM golang:1.24.4 AS backend_builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go run cmd/gentypes/main.go > /app/web/types.ts

FROM node AS client_builder
COPY --from=backend_builder /app /app
WORKDIR /app/web
RUN yarn
RUN yarn build

FROM golang:1.24.4
WORKDIR /app
COPY --from=client_builder /app /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend main.go
RUN go install github.com/mitranim/gow@latest
CMD ["/backend"]