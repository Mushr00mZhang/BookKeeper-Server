FROM golang:alpine AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct
COPY . .
# RUN apk add --no-cache gcc musl-dev
RUN go build -ldflags="-w -s" -o svc .

FROM scratch
WORKDIR /app
COPY --from=builder /svc .
EXPOSE 80
CMD ["/app/svc"]