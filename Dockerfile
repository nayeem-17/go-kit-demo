FROM golang:latest
LABEL maintainer="naim"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
# RUN go mod verify
# RUN go mod vendor
COPY . .
# ENV MONGO_URI=mongodb://root:example@mongo:27017/

RUN go build

CMD [ "./demo-go-kit" ]
