FROM node:lts-alpine3.24@sha256:21f403ab171f2dc89bad4dd69d7721bfd15f084ccb46cdd225f31f2bc59b5c9a

WORKDIR /app

COPY sass sass
COPY package.json package.json
COPY yarn.lock yarn.lock
COPY public public
# COPY sirius-header.js public/js/sirius-header.js

RUN yarn install --frozen-lockfile --ignore-scripts
RUN yarn compile-sass
ENTRYPOINT ["yarn", "serve"]