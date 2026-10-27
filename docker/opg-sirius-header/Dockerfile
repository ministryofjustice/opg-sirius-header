FROM node:lts-alpine3.24@sha256:1200d815de24f8dcca745fbf7651d166034a35adbcad8860a8538ffff44f1420

WORKDIR /app

COPY sass sass
COPY package.json .
COPY package-lock.json .
COPY public public
# COPY sirius-header.js public/js/sirius-header.js

RUN npm ci --ignore-scripts
RUN npm run compile-sass
ENTRYPOINT ["npm", "run", "serve"]
