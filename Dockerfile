FROM golang:1.16-buster
WORKDIR /app
COPY . .

ENV DATA_DIR=/ruehmkorf-data

RUN mkdir /ruehmkorf-data
RUN go build -o /ruehmkorf.com

CMD ["/ruehmkorf.com"]