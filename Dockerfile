FROM golang:1.16-alpine

WORKDIR /go_sales_api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./
RUN ls
RUN go build -o /go_sales_api_app main.go

EXPOSE 8090:3030
# Run
CMD [ "/go_sales_api_app" ]


ENV    MYSQL_HOST: ${MYSQL_HOST}
ENV    MYSQL_USER: ${MYSQL_USER}
ENV    MYSQL_PASSWORD: ${MYSQL_PASSWORD}
ENV    MYSQL_DBNAME: ${MYSQL_DBNAME}
ENV    JWT_SECRET: gotoko

#docker run -e MYSQL_HOST=host.docker.internal -e MYSQL_USER=aamir -e MYSQL_PASSWORD=aamir -e MYSQL_DBNAME=devcode_db -p 3030:3030 go_sales_api_app_v1