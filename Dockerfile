FROM quay.imanuel.dev/dockerhub/library---golang:1.16-alpine
WORKDIR /app
COPY . .

ENV DATA_DIR=/ruehmkorf-data

RUN mkdir /ruehmkorf-data
RUN go build -o /ruehmkorf.com

CMD ["/ruehmkorf.com"]