FROM golang:alpine AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .
# RUN apk add --no-cache gcc musl-dev
RUN go build -ldflags="-w -s" -o svc .

FROM scratch
WORKDIR /app
COPY --from=builder /app/svc .
EXPOSE 4000
CMD ["/app/svc"]