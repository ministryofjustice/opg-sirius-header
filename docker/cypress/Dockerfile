FROM cypress/included:16.1.1@sha256:e652cda84399aaf99c68d1d0322085dc3faa7ca9ace88d9bb4f0f045d2ce6764

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json .
COPY package-lock.json .

RUN npm ci --ignore-scripts

ENTRYPOINT ["npm", "run", "cypress-headless"]
