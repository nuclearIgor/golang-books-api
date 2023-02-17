FROM golang:1.19 as builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build
#ENTRYPOINT [ "app/server" ]

FROM scratch
COPY --from=builder /app/vueapi /vueapi
ENTRYPOINT [ "/vueapi" ]
