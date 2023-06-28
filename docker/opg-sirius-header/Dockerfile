FROM node:19.3.0-alpine3.16

COPY sass sass
COPY sirius-header.js sirius-header.js
COPY package.json package.json
COPY index.html index.html

RUN npm install
RUN npm run compile-sass
ENTRYPOINT ["npm", "run", "serve"]