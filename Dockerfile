FROM alpine:3.16.0
COPY hupi /
EXPOSE 8080
CMD ["./hupi", "start"]
