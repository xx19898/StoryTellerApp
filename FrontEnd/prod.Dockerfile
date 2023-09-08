FROM node:16.15.1 as build
WORKDIR /frontEnd
COPY package.json package-lock.json /frontEnd/
RUN npm ci
RUN mkdir /frontEnd/node_modules/.vite \
    && chown -R node:node /frontEnd/node_modules/.vite
COPY . /frontEnd/
RUN npm run build
FROM nginx:1.20.1
WORKDIR /etc/nginx/html
COPY --from=build /frontEnd/dist .



