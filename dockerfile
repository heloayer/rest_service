FROM golang:latest AS builder
WORKDIR /server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /server/cmd/main.go

FROM scratch
WORKDIR /bin/
COPY --from=builder /server/app .
ENV REDIS_ADDR redis:6379
ENV REDIS_PASSWORD ""
CMD [ "./app" ]
EXPOSE 8000


