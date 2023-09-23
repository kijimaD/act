###########
# builder #
###########

FROM golang:1.20-buster AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build \
      -ldflags='-w -s -extldflags "-static"' \
      -o ./bin/act \
 && upx-ucl --best --ultra-brute ./bin/act

###########
# release #
###########

FROM gcr.io/distroless/static-debian11:latest AS release
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    git
COPY --from=builder /build/bin/act /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/act"]
