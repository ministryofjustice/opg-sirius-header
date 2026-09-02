FROM cypress/included:15.21.1@sha256:8d874a8e5457459c7b8f185f797e46c299f7d0584c9fbda0b0ca23b3e85be749

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn install --frozen-lockfile --ignore-scripts

ENTRYPOINT ["yarn", "cypress-headless"]
