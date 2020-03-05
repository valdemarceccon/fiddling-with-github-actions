FROM golang:1.13-alpine as BUILDER

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server cmd/main.go  

FROM scratch

COPY --from=BUILDER /app/server /server

ENTRYPOINT [ "/server" ]