FROM golang:1.18-alpine as builder

ENV TZ=Asia/Tokyo
WORKDIR /go/src/app/

COPY ./app/ ./

RUN apk update \
    && apk add --no-cache git alpine-sdk tzdata \
    && go mod download \
    && go mod tidy \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./cmd/server/main ./cmd/server/

FROM gcr.io/distroless/static:nonroot as runner

WORKDIR /go/src/app/

# COPY --from=builder /go/src/app/static/* ./static/
# COPY --from=builder /go/src/app/main ./main
COPY --from=builder /go/src/app/infrastructure/public/* ./infrastructure/public/
COPY --from=builder /go/src/app/cmd/server/main ./cmd/server/main

ENTRYPOINT [ "./cmd/server/main" ]
