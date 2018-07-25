FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD  hello-world /bin/hello-world
ADD config.yml.dist /etc/hello-world/config.yml

CMD ["hello-world", "-config", "/etc/hello-world/config.yml"]