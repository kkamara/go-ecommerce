FROM golang:1.17-alpine

RUN apk --no-cache add make gcc g++

WORKDIR /app

COPY . ./

RUN go build -mod=readonly -o /ecommerce

CMD [ "/ecommerce" ]
