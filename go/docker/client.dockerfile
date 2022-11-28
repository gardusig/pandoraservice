FROM golang:1.19 AS base
WORKDIR /app
COPY go.mod ./
COPY integration_test.go ./

FROM base AS with_installed_package_version
ARG TEMPLATE_PACKAGE_VERSION
RUN go mod edit -require "github.com/gardusig/template-go@${TEMPLATE_PACKAGE_VERSION}"
RUN go mod tidy

FROM with_installed_package_version AS integration_tests
RUN go test .
