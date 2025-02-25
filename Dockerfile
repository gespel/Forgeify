FROM golang:1.24-bookworm
LABEL authors="stenh"

COPY . .
RUN go build
ENTRYPOINT ["./Forgeify"]