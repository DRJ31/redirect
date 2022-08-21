FROM alpine

COPY ./redirect /
RUN apk add --no-cache gcompat ca-certificates

ENTRYPOINT ["/redirect"]

EXPOSE 5000