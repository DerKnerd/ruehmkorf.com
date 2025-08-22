FROM docker.io/library/alpine:latest

COPY ruehmkorf.com /app/ruehmkorf.com

CMD ["/app/ruehmkorf.com", "serve"]
