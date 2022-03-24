FROM node:17.7.2
WORKDIR /app/react-ts
COPY package*.json ./
RUN yarn install
COPY . .
ENTRYPOINT [ "npm", "start" ]