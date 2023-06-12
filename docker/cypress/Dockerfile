FROM cypress/included:10.3.1

WORKDIR /root

RUN npm install "cypress-failed-log@2.9.2"

ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://sirius-header:8889

COPY cypress.config.js .
COPY cypress cypress