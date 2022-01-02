FROM golang:1.16-alpine3.15 as builder

RUN mkdir -p /build/src

ENV GOPATH=/build
ENV GOBIN=/usr/local/go/bin

ENV GO111MODULE=on
WORKDIR $GOPATH/src

COPY go.mod .
# COPY go.sum .

RUN go mod download
COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o out . \
    && cp out $GOPATH/.

CMD [ "/build/out" ]

FROM alpine:3.15
COPY --from=builder /build/out .
# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && chown appuser:appgroup out

USER appuser
CMD [ "./out" ]
