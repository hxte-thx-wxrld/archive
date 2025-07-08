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
COPY --from=client_builder /app /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/archive/backend main.go
RUN go install github.com/mitranim/gow@latest
RUN apt update && apt install -y python3-pip
#RUN python3 -m venv /opt/archive/py
#RUN source /opt/archive/py/bin/activate
RUN pip3 install --break-system-packages -r /app/requirements.daemon.txt
CMD ["/opt/archive/backend"]