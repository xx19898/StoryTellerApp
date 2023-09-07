FROM node:16.15.1 as INITIAL
WORKDIR /frontEnd
COPY package.json /frontEnd/
RUN npm install

RUN mkdir /frontEnd/node_modules/.vite \
    && chown -R node:node /frontEnd/node_modules/.vite

EXPOSE 3000

CMD ["npm","run","dev"]