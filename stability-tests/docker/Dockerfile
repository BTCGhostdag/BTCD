ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as BTCD
FROM ${KASPAMINER_IMAGE} as BTCDminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/BTCGhostdag/BTCD/

WORKDIR /go/src/github.com/BTCGhostdag/BTCD/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=BTCD /app/ /app/
COPY --from=BTCDminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/BTCGhostdag/BTCD/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]