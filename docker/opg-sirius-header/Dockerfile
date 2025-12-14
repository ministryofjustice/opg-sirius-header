FROM node:24.12.0-alpine3.22

WORKDIR /app

COPY sass sass
COPY package.json package.json
COPY public public
# COPY sirius-header.js public/js/sirius-header.js

RUN yarn install
RUN yarn compile-sass
ENTRYPOINT ["yarn", "serve"]