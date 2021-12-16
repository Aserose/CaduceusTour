FROM golang:1.17-alpine3.15 AS builder

COPY . /github.com/Aserose/CaduceusTour
WORKDIR /github.com/Aserose/CaduceusTour

RUN go mod download
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Aserose/CaduceusTour/bin/app .
COPY --from=0 /github.com/Aserose/CaduceusTour/internal/config config/
COPY --from=0 /github.com/Aserose/CaduceusTour/configs configs/

EXPOSE 3000

CMD ["./app"]