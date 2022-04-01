FROM golang:1.17-alpine
WORKDIR /opt/app/api

RUN apk update
RUN apk add git
RUN apk add openssh
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
ENV GO_PRIVATE="github.com/patrickchagastavares"
CMD [ "go", "run", "main.go" ]