# syntax=docker/dockerfile:1

FROM cgr.dev/chainguard/wolfi-base AS build

RUN apk upgrade --no-cache && apk add --no-cache go-1.26 ca-certificates

ARG PROVIDER_DIR
WORKDIR /src
COPY . .

RUN test -n "${PROVIDER_DIR}" \
    && test -f "${PROVIDER_DIR}/go.mod"

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    cd "${PROVIDER_DIR}" \
    && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/bin/provider .

RUN mkdir -p "/out/obot-tools/${PROVIDER_DIR}/bin" \
    && cp /out/bin/provider "/out/obot-tools/${PROVIDER_DIR}/bin/provider" \
    && if [ -d auth-providers-common/templates ]; then \
        mkdir -p /out/obot-tools/auth-providers-common/templates; \
        cp -R auth-providers-common/templates/. /out/obot-tools/auth-providers-common/templates/; \
    fi

FROM cgr.dev/chainguard/wolfi-base

RUN apk upgrade --no-cache && apk add --no-cache ca-certificates

ARG PROVIDER_DIR
ENV PORT=8000

COPY --from=build /out/obot-tools /obot-tools

EXPOSE 8000 9999
ENTRYPOINT ["/bin/provider"]
