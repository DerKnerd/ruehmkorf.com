ARG CI_DEPENDENCY_PROXY_GROUP_IMAGE_PREFIX

FROM ${CI_DEPENDENCY_PROXY_GROUP_IMAGE_PREFIX}/library/alpine:latest

COPY ruehmkorf.com /app/ruehmkorf.com

CMD ["/app/ruehmkorf.com", "serve"]
