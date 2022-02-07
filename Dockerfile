FROM golang:1.17-alpine


COPY . /app
WORKDIR /app
RUN go build
EXPOSE 8080
CMD ["./robo9"]
