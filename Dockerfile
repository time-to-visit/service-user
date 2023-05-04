FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .


RUN cd cmd  && go  build -o /service-user

EXPOSE 3001

CMD [ "/service-user" ]