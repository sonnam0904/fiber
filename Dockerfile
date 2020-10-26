FROM golang:1.14

ENV GO111MODULE=on
RUN go get -v github.com/gofiber/fiber/v2

RUN go get -v github.com/davecgh/go-spew/spew
RUN go get -v github.com/go-redis/redis
RUN go get -v gorm.io/driver/mysql
RUN go get -v gorm.io/gorm
RUN go get -u -v github.com/cosmtrek/air
RUN go get -v github.com/json-iterator/go
RUN go get -v go.mongodb.org/mongo-driver/mongo
RUN go get -v go.mongodb.org/mongo-driver/bson
RUN go get -v go.mongodb.org/mongo-driver/bson/primitive
RUN go get -v github.com/gofiber/template/html

ENV GO111MODULE=off
RUN go get -v github.com/golang/dep/cmd/dep

WORKDIR /go/src/fiber
EXPOSE 3000

CMD ["air"]

