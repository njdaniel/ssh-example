FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM sratch
COPY --from=builder /app/ssh-example /app/
ENTRYPOINT ["/app/ssh-example"]