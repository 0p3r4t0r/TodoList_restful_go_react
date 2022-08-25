FROM golang:1.19.0

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY ./* ./
RUN go get

EXPOSE $PORT

CMD [ "go", "run", "." ]