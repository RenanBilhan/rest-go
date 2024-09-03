FROM ubuntu:latest
LABEL authors="renan.bilhan"

ENTRYPOINT ["top", "-b"]