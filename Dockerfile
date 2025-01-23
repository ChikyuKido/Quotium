FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache gcc musl-dev upx nodejs npm

RUN npm install -g uncss

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY . .

RUN mkdir temp
RUN find . -name "*.html" -exec cp {} temp \;
RUN cp -R external/wat/website/css temp/
RUN uncss temp/*.html > external/wat/website/css/bulma.css
RUN rm -rf temp

COPY go.mod go.sum ./
RUN go mod download

RUN go build -ldflags="-s -w" Quotium

RUN upx --best --lzma Quotium

FROM alpine:latest

COPY --from=builder /app/Quotium /app/Quotium
COPY --from=builder /app/external/wat/website /app/external/wat/website
COPY --from=builder /app/website /app/website
WORKDIR /app
ENTRYPOINT ["./Quotium"]
