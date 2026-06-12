FROM cypress/included:15.17.0

WORKDIR /root

COPY package.json .
COPY yarn.lock .
RUN yarn install --frozen-lockfile --ignore-scripts

ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
ENTRYPOINT ["yarn", "cypress-headless"]