FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o school-prototype ./cmd/school-prototype/

FROM scratch AS final

COPY --from=builder /app/school-prototype /school-prototype

ENTRYPOINT ["/school-prototype"]

EXPOSE 8080