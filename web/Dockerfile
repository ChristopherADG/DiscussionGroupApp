FROM node:10.15.1-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json /app/package.json
RUN npm install

RUN npm install -g @vue/cli

# start app
CMD ["npm", "run", "serve"]