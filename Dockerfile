FROM golang:1.20.2-alpine3.17
WORKDIR /app-docker
EXPOSE 8080
COPY  . ./
RUN go build -o . .
CMD [ "./go-todoList" ]