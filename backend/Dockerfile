FROM golang

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY ./* ./
RUN go get

EXPOSE $PORT

CMD [ "go", "run", "." ]