FROM cypress/included:16.0.0@sha256:2ad59498de089a8eb03b6a3923c4e2f0801099f9be044ae2825531cb585e9370

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json .
COPY package-lock.json .

RUN npm ci --ignore-scripts

ENTRYPOINT ["npm", "run", "cypress-headless"]
