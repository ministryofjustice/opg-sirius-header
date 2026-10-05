FROM cypress/included:16.1.0@sha256:9618e4f3798bfc5434aeab1283813653437bbacc5d15922433b37f10b33aad49

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json .
COPY package-lock.json .

RUN npm ci --ignore-scripts

ENTRYPOINT ["npm", "run", "cypress-headless"]
