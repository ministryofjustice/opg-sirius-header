FROM cypress/included:15.17.0@sha256:9378699679f7d6ddb63c572866d4f11ec822abacb099345e89f5b86381efd53f

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json package.json
COPY yarn.lock yarn.lock
RUN yarn install --frozen-lockfile --ignore-scripts

ENTRYPOINT ["yarn", "cypress-headless"]