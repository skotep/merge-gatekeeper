ARG GO_VERSION=1.19.1
ARG IMAGE_REGISTRY=docker.io
FROM ${IMAGE_REGISTRY}/golang:${GO_VERSION}-alpine

ARG ORG=skotep
ARG REPO=merge-gatekeeper

ENV GO111MODULE=on LANG=en_US.UTF-8

RUN mkdir -p $GOPATH/src

WORKDIR ${GOPATH}/src/github.com/${ORG}/${REPO}

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/merge-gatekeeper \
    && mv merge-gatekeeper /go/bin/

ENTRYPOINT ["/go/bin/merge-gatekeeper"]
