FROM node:10.16-stretch AS builder-node

WORKDIR /src

COPY package.json yarn.lock ./
RUN yarn

COPY frontend ./frontend/
RUN yarn run build


FROM golang:1.12-alpine AS builder-go

WORKDIR /go/src/gitlab.com/totakoko/onetimesecret
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-extldflags '-static'"


FROM alpine:latest

COPY --from=builder-node /src/.build /opt/ots/.build
COPY --from=builder-go /go/src/gitlab.com/totakoko/onetimesecret/onetimesecret /opt/ots/onetimesecret

WORKDIR /opt/ots/
EXPOSE 5000
CMD ["/opt/ots/onetimesecret"]
