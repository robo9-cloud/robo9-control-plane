FROM golang:1.17-alpine


COPY . /app
WORKDIR /app
RUN go mod tidy && go build
EXPOSE 18080
CMD ["./robo9-cli","server"]
