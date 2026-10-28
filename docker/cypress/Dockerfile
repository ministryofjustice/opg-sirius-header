FROM cypress/included:16.2.0@sha256:c196ed896aec24c7ea8fc1d4185b210493e311339b6c09d0153f9b8dfc6a6058

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json .
COPY package-lock.json .

RUN npm ci --ignore-scripts

ENTRYPOINT ["npm", "run", "cypress-headless"]
