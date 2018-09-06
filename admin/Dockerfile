FROM node:8

ENV HOST 0.0.0.0
ENV PORT=3010
ENV BASE_URL=https://odptapi.theliveup.tv
EXPOSE 3010

RUN mkdir -p /app
COPY . /app
WORKDIR /app

RUN yarn install && \
    yarn build

CMD ["yarn", "start"]
