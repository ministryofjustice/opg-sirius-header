FROM cypress/included:16.3.0@sha256:3efa8d156a094a6c13ddf3fba616e97f7492706685064b75a39823e1147eca97

WORKDIR /root


ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://opg-sirius-header:8080

COPY cypress.config.js .
COPY cypress cypress
COPY package.json .
COPY package-lock.json .

RUN npm ci --ignore-scripts

ENTRYPOINT ["npm", "run", "cypress-headless"]
