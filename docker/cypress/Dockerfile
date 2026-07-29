FROM cypress/included:15.19.0@sha256:fe59f2cac5c7451644c8780cb2fb6735fde512d06642117199a7479269702f03

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn install --frozen-lockfile --ignore-scripts

ENTRYPOINT ["yarn", "cypress-headless"]
