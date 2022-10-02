###########
# builder #
###########

FROM golang:1.13-buster AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build \
      -ldflags='-w -s -extldflags "-static"' \
      -o ./bin/stats \
 && upx-ucl --best --ultra-brute ./bin/stats

###########
# release #
###########

FROM golang:1.13-buster AS release
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    git

COPY --from=builder /build/bin/stats /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/stats"]
