FROM alpine:latest

COPY onetimesecret /usr/local/bin/onetimesecret

EXPOSE 5000
CMD ["onetimesecret"]
