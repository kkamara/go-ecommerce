FROM golang:1.17-alpine

RUN apk --no-cache add make gcc g++

ENV GO111MODULE=on
ENV GOMODCACHE=/tmp

WORKDIR /app

COPY . ./

RUN go build -mod=mod -o /ecommerce

CMD [ "/ecommerce" ]
