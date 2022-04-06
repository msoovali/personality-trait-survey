### BUILD ###
FROM golang:1.18-alpine AS BUILD
RUN apk add git
WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go test ./... -cover
RUN go build -o ./out/trait-survey ./cmd/api/main.go

### DEPLOY ###
FROM alpine:3.15
RUN apk add ca-certificates

COPY --from=BUILD /tmp/app/out/trait-survey /app/trait-survey
COPY --from=BUILD /tmp/app/conf /app/conf

WORKDIR app
EXPOSE 8080
# Run the binary program produced by `go build`
CMD ["./trait-survey"]