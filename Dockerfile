FROM docker.io/library/node:alpine AS build-frontend
WORKDIR /app

COPY . .

WORKDIR /app/public/admin

RUN npm install

FROM docker.io/library/golang:1.24-alpine AS build-backend
WORKDIR /app

COPY . .

RUN go build -o ruehmkorf.com

FROM docker.io/library/alpine:latest
COPY --from=build-frontend /app/public /app/public
COPY --from=build-backend /app/frontend/templates /app/frontend/templates
COPY --from=build-backend /app/admin/templates /app/admin/templates
COPY --from=build-backend /app/ruehmkorf.com /app/ruehmkorf.com

WORKDIR /app

ENV DATA_DIR=/ruehmkorf-data

RUN mkdir /ruehmkorf-data

CMD ["/app/ruehmkorf.com"]
