FROM node:10.8-stretch AS builder-node

WORKDIR /src

COPY package.json yarn.lock ./
RUN yarn

COPY gulpfile.js .
COPY templates ./templates/
COPY styles ./styles/
RUN yarn run gulp build


FROM golang:1.10-alpine AS builder-go

WORKDIR /go/src/gitlab.dreau.fr/home/onetimesecret
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-extldflags '-static'"


FROM alpine:latest

COPY --from=builder-node /src/.build /opt/ots/.build
COPY --from=builder-go /go/src/gitlab.dreau.fr/home/onetimesecret/onetimesecret /opt/ots/onetimesecret

WORKDIR /opt/ots/
EXPOSE 5000
CMD ["/opt/ots/onetimesecret"]
