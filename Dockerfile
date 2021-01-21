FROM alpine:latest
ARG BINARY=api-server
COPY ./conf/  /api-server/conf/
COPY ./bin/$BINARY  /api-server/api-server
RUN chmod 755 /api-server/api-server
EXPOSE 9001
ENTRYPOINT ["/api-server/api-server", "-c", "/api-server/conf/config.yaml"]
