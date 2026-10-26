FROM node:22.23.4-alpine3.24@sha256:025b8e975a61d94e33abeafaad1255ef44e4898cf1a1b6446860009c72e6d284

WORKDIR /app

COPY sass sass
COPY package.json .
COPY package-lock.json .
COPY public public
# COPY sirius-header.js public/js/sirius-header.js

RUN npm ci --ignore-scripts
RUN npm run compile-sass
ENTRYPOINT ["npm", "run", "serve"]
