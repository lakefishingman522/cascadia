# syntax=docker/dockerfile:1
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"
ARG GO_VERSION="1.20"

# Builder image
FROM golang:1.20-alpine3.18 AS build-env

ARG GIT_VERSION
ARG GIT_COMMIT

RUN set -eux; \
    apk add --no-cache \
    ca-certificates=20230506-r0 \
    build-base=0.5-r3 \
    linux-headers=6.3-r0 \
    git=2.40.1-r0 \
    psmisc \
    openssh \
    bash=5.2.15-r5

RUN mkdir /cascadia
WORKDIR /cascadia

RUN mkdir -p /root/.ssh
RUN chmod -R 600 /root/.ssh/
RUN ssh-keyscan github.com >>/root/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN go env -w GOPRIVATE=github.com/cascadia-protocol/*

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=secret,id=sshKey,dst=/root/.ssh/id_ecdsa \
    go mod download

RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
    -O /lib/libwasmvm_muslc.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1)

# Add source code
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    # --mount=type=bind,source=.,target=/utility \
    GOWORK=off go build \
        -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="cascadia" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="cascadiad" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o /cascadia/build/cascadiad \
        ./cmd/cascadiad

# Build Runner

FROM ${RUNNER_IMAGE}

COPY --from=build-env /cascadia/build/cascadiad /bin/cascadiad
ENV HOME /cascadia
WORKDIR $HOME
EXPOSE 26656 26657 1317 1318 9090 9091

ENTRYPOINT ["cascadiad"]
