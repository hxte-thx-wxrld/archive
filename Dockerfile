FROM node AS client_builder
WORKDIR /app
COPY . .
WORKDIR /app/web
RUN yarn
RUN yarn build
FROM golang:1.24.4
COPY --from=client_builder /app /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server main.go
RUN go install github.com/mitranim/gow@latest
CMD ["/server"]