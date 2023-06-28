WORKDIR /app

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

FROM node:19.3.0-alpine3.16 as asset-env

WORKDIR /app

COPY sass sass
COPY sirius-header.js sirius-header.js
COPY package.json .

RUN npm install
RUN npm run build

FROM alpine:3

RUN apk --update --no-cache add \
    ca-certificates \
    tzdata

# Patch vulnerabilities
RUN apk upgrade --no-cache libcrypto3 libssl3

COPY --from=asset-env /app/index.html index.html
COPY --from=asset-env /app/sirius-header.js sirius-header.js
COPY --from=asset-env /app/sass sass
ENTRYPOINT ["./opg-sirius-workflow"]
