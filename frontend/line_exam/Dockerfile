# syntax=docker/dockerfile:1

FROM node:16.5.0-alpine

WORKDIR /app

#add /app/node_modules/.bin to $PATH
ENV PATH /app/node_modules/.bin:$PATH

#install and cache app dependencies
COPY package.json /app/package.json
RUN yarn install
RUN yarn global add @vue/cli

#start app
CMD ["yarn", "run", "serve"]


#run docker command
#docker build --tag docker-frontend-vuejs
#docker run -v ${PWD}:/app -v /app/node_modules -p 8081:8080 --rm docker-frontend-vuejs

