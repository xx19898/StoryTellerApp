FROM node:16.15.1 as INITIAL
WORKDIR /frontEnd
COPY package.json /frontEnd/
RUN npm install
COPY . /frontEnd/
EXPOSE 3000
CMD ["npm","run","dev"]