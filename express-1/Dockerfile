FROM ghcr.io/hongfs/env:node20

WORKDIR  /code

COPY . .

RUN npm install

EXPOSE 9000

CMD [ "node", "index.js"]
