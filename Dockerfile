FROM ubuntu:latest
LABEL authors="stenh"

ENTRYPOINT ["top", "-b"]