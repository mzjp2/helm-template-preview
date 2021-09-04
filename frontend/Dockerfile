FROM node:lts-alpine

COPY package-lock.json package.json /app/
WORKDIR /app

RUN npm ci

COPY . /app/

EXPOSE 8080
CMD ["npm", "start"]
