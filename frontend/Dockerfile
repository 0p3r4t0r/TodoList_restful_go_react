FROM node:16.17.0-alpine

RUN npm install -g pnpm

ENV NODE_ENV=development
ENV PORT=3000

WORKDIR /app

COPY ./ ./
RUN pnpm install

EXPOSE $PORT

CMD [ "pnpm", "start" ]