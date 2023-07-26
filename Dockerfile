FROM quay.imanuel.dev/dockerhub/library---node:latest AS build-frontend
WORKDIR /app
COPY . .
RUN cd public/admin && npm install

FROM quay.imanuel.dev/dockerhub/library---golang:1.20-alpine
WORKDIR /app
COPY . .
COPY --from=build-frontend /app/public /app/public

ENV DATA_DIR=/ruehmkorf-data

RUN go build -o ruehmkorf.com
RUN mkdir /ruehmkorf-data
RUN rm /app/public/admin/package.json /app/public/admin/package-lock.json

CMD ["/app/ruehmkorf.com"]