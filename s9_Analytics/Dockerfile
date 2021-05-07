FROM node:12.22.1-alpine
WORKDIR /usr/app/
RUN npm i -g nodemon
COPY ./package.json .
COPY . .
RUN npm install
EXPOSE 8080
CMD npm start
